package lisp

import (
	"unicode"
)

type Pattern interface {
	Match(Value) bool
}

type symbolPattern struct {
	symbol Value
}

func (p *symbolPattern) Match(v Value) bool {
	return p.symbol.IsEq(v)
}

type variablePattern struct {
	name Value
}

func (p *variablePattern) Match(Value) bool {
	return true
}

type pairPattern struct {
	head Pattern
	tail Pattern
}

func (p *pairPattern) Match(v Value) bool {
	if v.IsAtom() {
		return false
	} else {
		return p.head.Match(v.GetCar()) && p.tail.Match(v.GetCdr())
	}
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

// utilities
func isUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
