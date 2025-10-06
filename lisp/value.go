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
type cons struct {
	car Value
	cdr Value
}

func Cons(car Value, cdr Value) Value {
	return &cons{car, cdr}
}

func List(e ...Value) (val Value) {
	val = Nil
	for i := len(e) - 1; i >= 0; i-- {
		val = Cons(e[i], val)
	}
	return
}

func (*cons) IsAtom() bool {
	return false
}

func (c *cons) GetCar() Value {
	return c.car
}

func (c *cons) GetCdr() Value {
	return c.cdr
}
func (c *cons) IsEq(other Value) bool {
	return false
}

func (c *cons) String() string {
	var sb strings.Builder

	cdr := c.GetCdr()
	sb.WriteString("(")
	sb.WriteString(c.GetCar().String())
	for {
		if cdr.IsAtom() {
			break
		}
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
