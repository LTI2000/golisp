package lisp

func Eval(e Expression) (Value, error) {
	return e.Eval()
}

func (l *literal) Eval() (Value, error) {
	return QUOTE(l.value), nil
}

func (p *prim_app1) Eval() (Value, error) {
	if x, err := p.arg0.Eval(); err != nil {
		return nil, err
	} else {
		return p.prim(x), nil
	}
}

func (p *prim_app2) Eval() (Value, error) {
	if x, err := p.arg0.Eval(); err != nil {
		return nil, err
	} else if y, err := p.arg1.Eval(); err != nil {
		return nil, err
	} else {
		return p.prim(x, y), nil
	}
}
