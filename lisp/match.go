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
	index  Index
}

type Index = []idx
type idx struct{ i, n int }

func NewBindings() *Bindings {
	return &Bindings{}
}

func (b *Bindings) Extend(name string, value Value, index Index) *Bindings {
	return &Bindings{b, name, value, index}
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
		if b.index != nil {
			sb.WriteString(" @")
			sb.WriteString(fmt.Sprintf("%v", b.index))
		}
		sb.WriteString("]")
		b = b.parent
	}
	sb.WriteString("}")
	return sb.String()
}

// Pattern
type Pattern interface {
	Match(*Bindings, Value, Index) (*Bindings, bool)
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

func (p *symbolPattern) Match(b *Bindings, v Value, i Index) (*Bindings, bool) {
	return b, p.symbol.IsEq(v)
}

// variablePattern
type variablePattern struct {
	name string
}

func (p *variablePattern) Match(b *Bindings, v Value, i Index) (*Bindings, bool) {
	return b.Extend(p.name, v, i), true
}

// pairPattern
type pairPattern struct {
	head Pattern
	tail Pattern
}

func (p *pairPattern) Match(b *Bindings, v Value, i Index) (*Bindings, bool) {
	if v.IsAtom() {
		return b, false
	} else {
		if b1, matches := p.head.Match(b, v.GetCar(), i); matches {
			if b2, matches := p.tail.Match(b1, v.GetCdr(), i); matches {
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

func (p *repeatingPattern) Match(b *Bindings, value Value, index Index) (*Bindings, bool) {
	vs := Slice(value)
	n := len(vs)
	for i, v := range vs {
		b1, matches := p.pattern.Match(b, v, append(index, idx{i, n}))
		if matches {
			b = b1
		} else {
			return nil, false
		}
		i++
	}
	return b, true
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
