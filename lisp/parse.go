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
	if bindings, matches := quote.Match(NewBindings(), value); matches {
		if value, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else {
			return &literal{value}, nil
		}
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
	String() string
}

// literal
type literal struct {
	value Value
}

func (l *literal) String() string {
	return "(quote " + l.value.String() + ")"
}

// prim_app1
type prim_app1 struct {
	prim func(Value) Value
	arg0 Expression
}

func (p *prim_app1) String() string {
	panic("unimplemented")
}

// prim_app2
type prim_app2 struct {
	prim func(Value, Value) Value
	arg0 Expression
	arg1 Expression
}

func (p *prim_app2) String() string {
	panic("unimplemented")
}
