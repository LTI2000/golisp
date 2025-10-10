package lisp

import "strings"

// Expression
type Expression interface {
	String() string
	Eval() (Value, error)
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

// cond

type clause struct {
	predicate  Expression
	expression Expression
}

type conditonal struct {
	clauses []clause
}

func (l *conditonal) String() string {
	var sb strings.Builder

	sb.WriteString("(cond")
	for _, clause := range l.clauses {
		sb.WriteString(" (")
		sb.WriteString(clause.predicate.String())
		sb.WriteString(" ")
		sb.WriteString(clause.expression.String())
		sb.WriteString(")")
	}
	sb.WriteString(")")

	return sb.String()
}
