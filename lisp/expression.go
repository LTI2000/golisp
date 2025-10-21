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
var ATOM Expression = Symbol("atom")
var EQ Expression = Symbol("eq")
var CAR Expression = Symbol("car")
var CDR Expression = Symbol("cdr")
var CONS Expression = Symbol("cons")
var COND Expression = Symbol("cond")
var LABEL Expression = Symbol("label")
var LAMBDA Expression = Symbol("lambda")
var DEFUN Expression = Symbol("defun")

// Symbol
type symbol struct {
	name string
}

var symbols map[string]*symbol = make(map[string]*symbol)

func Symbol(name string) Expression {
	expression, ok := symbols[name]
	if !ok {
		expression = &symbol{name}
		symbols[name] = expression
	}
	return expression
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

func Uncons(e Expression) (Expression, Expression, error) {
	switch c := e.(type) {
	case *cons:
		return c.car, c.cdr, nil
	default:
		return nil, nil, fmt.Errorf("Uncons: not a cons: %v", c)
	}
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

// create a (possibly empty) List from a Value slice.
func List(slice ...Expression) Expression {
	list := NIL
	for i := len(slice) - 1; i >= 0; i-- {
		list = Cons(slice[i], list)
	}
	return list
}

func Bool(b bool) Expression {
	if b {
		return T
	} else {
		return NIL
	}
}
