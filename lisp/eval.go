package lisp

func Eval(e Expression) (Value, error) {
	return e.Eval()
}

func (e *literal) Eval() (Value, error) {
	return e.value, nil
}

func (e *prim_app1) Eval() (Value, error) {
	panic("NYI")
}

func (e *prim_app2) Eval() (Value, error) {
	panic("NYI")
}
