package lisp

import (
	"errors"
	"strings"
	"unicode"
)

// Bindings
type Bindings struct {
	name  string
	value Value
	next  *Bindings
}

func extend(b *Bindings, name string, value Value) *Bindings {
	return &Bindings{name, value, b} // FIXME add check for duplicate names
}

func extendList(b *Bindings, name string, value Value) *Bindings {
	if b != nil {
		if b.name != name {
			return &Bindings{b.name, b.value, extendList(b.next, name, value)}
		} else {
			return &Bindings{b.name, Append(b.value, List(value)), b.next}
		}
	} else {
		return &Bindings{name, List(value), b}
	}
}

func Lookup(b *Bindings, name string) (Value, error) {
	if b == nil {
		return nil, errors.New(name + " is unbound")
	} else if name == b.name {
		return b.value, nil
	} else {
		return Lookup(b.next, name)
	}
}

func String(b *Bindings) string {
	var sb strings.Builder

	sb.WriteString("{")
	for b != nil {
		sb.WriteString("[")
		sb.WriteString(b.name)
		sb.WriteString("=")
		sb.WriteString(b.value.String())
		sb.WriteString("]")
		b = b.next
	}
	sb.WriteString("}")
	return sb.String()
}

// Pattern
type Pattern interface {
	Match(*Bindings, Value, bool) (*Bindings, bool)
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

func (p *symbolPattern) Match(b *Bindings, v Value, i bool) (*Bindings, bool) {
	return b, p.symbol.IsEq(v)
}

// variablePattern
type variablePattern struct {
	name string
}

func (p *variablePattern) Match(b *Bindings, v Value, i bool) (*Bindings, bool) {
	if i {
		return extendList(b, p.name, v), true
	} else {
		return extend(b, p.name, v), true
	}
}

// pairPattern
type pairPattern struct {
	head Pattern
	tail Pattern
}

func (p *pairPattern) Match(b *Bindings, v Value, i bool) (*Bindings, bool) {
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

func (p *repeatingPattern) Match(b *Bindings, v Value, i bool) (*Bindings, bool) {
	for !v.IsAtom() {
		var matches bool
		b, matches = p.pattern.Match(b, v.GetCar(), true)
		if matches {
			v = v.GetCdr()
		} else {
			return nil, false
		}
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
