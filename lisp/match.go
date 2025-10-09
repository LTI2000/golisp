package lisp

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Bindings
type Bindings struct {
	parent *Bindings
	name   string
	value  Value
}

func NewBindings() *Bindings {
	return &Bindings{}
}

func (b *Bindings) Extend(name string, value Value) *Bindings {
	return &Bindings{b, name, value}
}

func (b *Bindings) Lookup(name string) (Value, error) {
	if name == b.name {
		return b.value, nil
	} else if b.parent != nil {
		return b.parent.Lookup(name)
	} else {
		return nil, errors.New(name + " is unbound")
	}
}

func (b *Bindings) String() string {
	var sb strings.Builder

	sb.WriteString("{")
	for b.parent != nil {
		sb.WriteString("[")
		sb.WriteString(b.name)
		sb.WriteString("=")
		sb.WriteString(b.value.String())
		sb.WriteString("]")
		b = b.parent
	}
	sb.WriteString("}")
	return sb.String()
}

// Pattern
type Pattern interface {
	Match(*Bindings, Value) (*Bindings, bool)
}

func NewPattern(pattern Value) Pattern {
	if pattern.IsAtom() {
		name := pattern.String()
		if isUpperCase(name) {
			return &variablePattern{name}
		} else {
			return &symbolPattern{pattern}
		}
	} else {
		head, tail := pattern.GetCar(), pattern.GetCdr()
		if tail.String() == "(...)" { // FIXME crude check
			return &repeatingPattern{NewPattern(head)}
		} else {
			return &pairPattern{NewPattern(head), NewPattern(tail)}
		}
	}
}

// symbolPattern
type symbolPattern struct {
	symbol Value
}

func (p *symbolPattern) Match(b *Bindings, v Value) (*Bindings, bool) {
	return b, p.symbol.IsEq(v)
}

// variablePattern
type variablePattern struct {
	name string
}

func (p *variablePattern) Match(b *Bindings, v Value) (*Bindings, bool) {
	return b.Extend(p.name, v), true
}

// pairPattern
type pairPattern struct {
	head Pattern
	tail Pattern
}

func (p *pairPattern) Match(b *Bindings, v Value) (*Bindings, bool) {
	if v.IsAtom() {
		return b, false
	} else {
		if b1, matches := p.head.Match(b, v.GetCar()); matches {
			if b2, matches := p.tail.Match(b1, v.GetCdr()); matches {
				return b2, true
			}
		}
		return b, false
	}
}

// repeatingPattern
type repeatingPattern struct {
	pattern Pattern
}

func (p *repeatingPattern) Match(b *Bindings, v Value) (*Bindings, bool) {
	fmt.Println("REPEAT MATCH")
	for !v.IsAtom() {
		head := v.GetCar()
		v = v.GetCdr()
		fmt.Println("HEAD " + head.String())
		_, matches := p.pattern.Match(b, head)
		if matches {
			fmt.Println("MATCH " + head.String())
		} else {
			return nil, false
		}
	}
	return b, false
}

// utilities
func isUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
