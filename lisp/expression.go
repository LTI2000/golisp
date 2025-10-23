package lisp

import (
	"fmt"
	"strings"
)

// Expression
type Expression interface {
	prim_atom() bool
	prim_car() (Expression, error)
	prim_cdr() (Expression, error)
	prim_eq(Expression) bool
	String() string
}

func Atom(x Expression) bool {
	return x.prim_atom()
}

func Car(x Expression) (Expression, error) {
	return x.prim_car()
}

func Cdr(x Expression) (Expression, error) {
	return x.prim_cdr()
}
func Eq(x, y Expression) bool {
	return x.prim_eq(y)
}

var T Expression = Symbol("t")
var NIL Expression = Symbol("nil")
var QUOTE Expression = Symbol("quote")
var LABEL Expression = Symbol("label")
var LAMBDA Expression = Symbol("lambda")

// Symbol
type symbol struct {
	name string
}

var symbols map[string]*symbol = make(map[string]*symbol)

func Symbol(name string) Expression {
	s, ok := symbols[name]
	if !ok {
		s = &symbol{name}
		symbols[name] = s
	}
	return s
}

func (*symbol) prim_atom() bool {
	return true
}

func (s *symbol) prim_car() (Expression, error) {
	return nil, fmt.Errorf("car: not a cons: %v", s)
}

func (s *symbol) prim_cdr() (Expression, error) {
	return nil, fmt.Errorf("cdr: not a cons: %v", s)
}

func (s *symbol) prim_eq(other Expression) bool {
	switch e := other.(type) {
	case *symbol:
		return s.name == e.name
	default:
		return false
	}
}

func (s *symbol) String() string {
	return s.name
}

// Pair
type cons struct {
	car Expression
	cdr Expression
}

func Cons(car Expression, cdr Expression) Expression {
	return &cons{car, cdr}
}

func (*cons) prim_atom() bool {
	return false
}

func (c *cons) prim_car() (Expression, error) {
	return c.car, nil
}

func (c *cons) prim_cdr() (Expression, error) {
	return c.cdr, nil
}
func (c *cons) prim_eq(other Expression) bool {
	return false
}

func (c *cons) String() string {
	var sb strings.Builder
	sb.WriteString("(")

	sb.WriteString(c.car.String())
	rest := c.cdr
loop:
	for {
		switch e := rest.(type) {
		case *cons:
			sb.WriteString(" ")
			sb.WriteString(e.car.String())
			rest = e.cdr
		case *symbol:
			if e != NIL {
				sb.WriteString(" . ")
				sb.WriteString(e.String())
			}
			break loop
		}
	}

	sb.WriteString(")")
	return sb.String()
}

// Utils

// create a (possibly empty) list from an expression slice.
func List(slice ...Expression) Expression {
	list := NIL
	for i := len(slice) - 1; i >= 0; i-- {
		list = Cons(slice[i], list)
	}
	return list
}

// create a (possibly empty) expression slice from a list. Panics if list is not a proper list.
func Slice(list Expression) (slice []Expression) {
	return SliceMapped(list, Id)
}

// create a (possibly empty) slice from a list. Uses f to map elements to the target type T. Panics if list is not a proper list.
func SliceMapped[T any](list Expression, f func(e Expression) T) (slice []T) {
	for list != NIL {
		slice = append(slice, f(Must(Car, list)))
		list = Must(Cdr, list)
	}
	return
}

func Bool(b bool) Expression {
	if b {
		return T
	} else {
		return NIL
	}
}
