package lisp

import (
	"errors"
	"fmt"
)

var quote = NewPattern(Must(Read, "(quote X)"))
var atom = NewPattern(Must(Read, "(atom X)"))
var eq = NewPattern(Must(Read, "(eq X Y)"))
var car = NewPattern(Must(Read, "(car X)"))
var cdr = NewPattern(Must(Read, "(cdr X)"))
var cons = NewPattern(Must(Read, "(cons X Y)"))
var cond = NewPattern(Must(Read, "(cond (P E) ...)"))

func ParseExpression(value Value) (Expression, error) {
	if bindings, matches := quote.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else {
			return &literal{x}, nil
		}
	} else if bindings, matches := atom.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{ATOM, "atom", arg0}, nil
		}
	} else if bindings, matches := eq.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else if y, err := Lookup(bindings, "Y"); err != nil {
			return nil, err
		} else if arg1, err := ParseExpression(y); err != nil {
			return nil, err
		} else {
			return &prim_app2{EQ, "eq", arg0, arg1}, nil
		}
	} else if bindings, matches := car.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{CAR, "car", arg0}, nil
		}
	} else if bindings, matches := cdr.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{CDR, "cdr", arg0}, nil
		}
	} else if bindings, matches := cons.Match(nil, value, false); matches {
		if x, err := Lookup(bindings, "X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else if y, err := Lookup(bindings, "Y"); err != nil {
			return nil, err
		} else if arg1, err := ParseExpression(y); err != nil {
			return nil, err
		} else {
			return &prim_app2{CONS, "cons", arg0, arg1}, nil
		}
	} else if bindings, matches := cond.Match(nil, value, false); matches {
		fmt.Printf("BINDINGS: %v\n", bindings)
		panic("NYI")
	} else {
		return nil, errors.New("illegal expression: " + value.String())
	}
}
