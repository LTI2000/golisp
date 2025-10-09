package lisp

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
