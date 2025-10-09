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
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else {
			return &literal{x}, nil
		}
	} else if bindings, matches := atom.Match(NewBindings(), value); matches {
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{ATOM, "atom", arg0}, nil
		}
	} else if bindings, matches := eq.Match(NewBindings(), value); matches {
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else if y, err := bindings.Lookup("Y"); err != nil {
			return nil, err
		} else if arg1, err := ParseExpression(y); err != nil {
			return nil, err
		} else {
			return &prim_app2{EQ, "eq", arg0, arg1}, nil
		}
	} else if bindings, matches := car.Match(NewBindings(), value); matches {
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{CAR, "car", arg0}, nil
		}
	} else if bindings, matches := cdr.Match(NewBindings(), value); matches {
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else {
			return &prim_app1{CDR, "cdr", arg0}, nil
		}
	} else if bindings, matches := cons.Match(NewBindings(), value); matches {
		if x, err := bindings.Lookup("X"); err != nil {
			return nil, err
		} else if arg0, err := ParseExpression(x); err != nil {
			return nil, err
		} else if y, err := bindings.Lookup("Y"); err != nil {
			return nil, err
		} else if arg1, err := ParseExpression(y); err != nil {
			return nil, err
		} else {
			return &prim_app2{CONS, "cons", arg0, arg1}, nil
		}
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
	name string
	arg0 Expression
}

func (p *prim_app1) String() string {
	return "(" + p.name + " " + p.arg0.String() + ")"
}

// prim_app2
type prim_app2 struct {
	prim func(Value, Value) Value
	name string
	arg0 Expression
	arg1 Expression
}

func (p *prim_app2) String() string {
	return "(" + p.name + " " + p.arg0.String() + " " + p.arg1.String() + ")"
}
