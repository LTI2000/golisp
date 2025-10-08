package lisp

import (
	"strings"
	"unicode"
)

// Bindings
type Bindings struct {
	parent *Bindings
	name   Value
	value  Value
}

func NewBindings() *Bindings {
	return &Bindings{}
}

func (b *Bindings) Extend(name Value, value Value) *Bindings {
	return &Bindings{b, name, value}
}

func (b *Bindings) String() string {
	var sb strings.Builder

	sb.WriteString("{")
	for b.parent != nil {
		sb.WriteString("[")
		sb.WriteString(b.name.String())
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
		if isUpperCase(pattern.String()) {
			return &variablePattern{pattern}
		} else {
			return &symbolPattern{pattern}
		}
	} else {
		return &pairPattern{NewPattern(pattern.GetCar()), NewPattern(pattern.GetCdr())}
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
	name Value
}

func (p *variablePattern) Match(b *Bindings, v Value) (*Bindings, bool) {
	return b.Extend(p.name, v), true // FIXME extend bindings
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

// utilities
func isUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
