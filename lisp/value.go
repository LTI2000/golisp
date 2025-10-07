package lisp

import "strings"

// Value
type Value interface {
	IsAtom() bool
	GetCar() Value
	GetCdr() Value
	IsEq(other Value) bool
	String() string
}

var T Value = Symbol("t")
var Nil Value = Symbol("nil")
var Quote Value = Symbol("quote")

// Symbol
type symbol struct {
	name string
}

var symbols map[string]*symbol = make(map[string]*symbol)

func Symbol(name string) Value {
	value, ok := symbols[name]
	if !ok {
		value = &symbol{name}
		symbols[name] = value
	}
	return value
}

func BoolSymbol(b bool) Value {
	if b {
		return T
	} else {
		return Nil
	}
}

func (*symbol) IsAtom() bool {
	return true
}

func (*symbol) GetCar() Value {
	panic("GetCar(): got symbol")
}

func (*symbol) GetCdr() Value {
	panic("GetCdr(): got symbol")
}

func (s *symbol) IsEq(other Value) bool {
	switch v := other.(type) {
	case *symbol:
		return s.name == v.name
	default:
		return false
	}
}

func (s *symbol) String() string {
	if s == Nil {
		return "()"
	} else {
		return s.name
	}
}

// Cons
type pair struct {
	head Value
	tail Value
}

func Pair(head Value, tail Value) Value {
	return &pair{head, tail}
}

func (*pair) IsAtom() bool {
	return false
}

func (p *pair) GetCar() Value {
	return p.head
}

func (p *pair) GetCdr() Value {
	return p.tail
}
func (p *pair) IsEq(other Value) bool {
	return false
}

func (p *pair) String() string {
	var sb strings.Builder

	cdr := p.GetCdr()
	sb.WriteString("(")
	sb.WriteString(p.GetCar().String())
	for !cdr.IsAtom() {
		sb.WriteString(" ")
		sb.WriteString(cdr.GetCar().String())
		cdr = cdr.GetCdr()
	}
	if cdr != Nil {
		sb.WriteString(" . ")
		sb.WriteString(cdr.String())
	}
	sb.WriteString(")")

	return sb.String()
}

// Utils
func List(slice ...Value) Value {
	list := Nil
	for i := len(slice) - 1; i >= 0; i-- {
		list = Pair(slice[i], list)
	}
	return list
}

func Slice(list Value) (slice []Value) {
	for list != Nil {
		slice = append(slice, list.GetCar())
		list = list.GetCdr()
	}
	return
}
