package lisp

import "errors"

var quote = NewPattern(Must(Read, "(quote X)"))
var atom = NewPattern(Must(Read, "(atom X)"))
var eq = NewPattern(Must(Read, "(eq X Y)"))
var car = NewPattern(Must(Read, "(car X)"))
var cdr = NewPattern(Must(Read, "(cdr X)"))
var cons = NewPattern(Must(Read, "(cons X Y)"))
var cond = NewPattern(Must(Read, "(cond (P E) ...)"))

func ParseExpression(value Value) (Expression, error) {
	if _, matches := quote.Match(NewBindings(), value); matches {
		return &literal{nil}, nil
	} else if _, matches := atom.Match(NewBindings(), value); matches {
		return &prim_app1{ATOM, nil}, nil
	} else if _, matches := eq.Match(NewBindings(), value); matches {
		return &prim_app2{EQ, nil, nil}, nil
	} else if _, matches := car.Match(NewBindings(), value); matches {
		return &prim_app1{CAR, nil}, nil
	} else if _, matches := cdr.Match(NewBindings(), value); matches {
		return &prim_app1{CDR, nil}, nil
	} else if _, matches := cons.Match(NewBindings(), value); matches {
		return &prim_app2{CONS, nil, nil}, nil
	} else {
		return nil, errors.New("illegal expression")
	}
}

// Expression
type Expression interface {
}

type literal = struct {
	value Value
}

type prim_app1 = struct {
	prim func(Value) Value
	arg0 Value
}

type prim_app2 = struct {
	prim func(Value, Value) Value
	arg0 Value
	arg1 Value
}
