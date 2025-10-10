package lisp

import (
	"errors"
)

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

func (c *conditonal) Eval() (Value, error) {
	for _, clause := range c.clauses {
		if pval, err := clause.predicate.Eval(); err != nil {
			return nil, err
		} else if pval.IsEq(T) {
			if val, err := clause.expression.Eval(); err != nil {
				return nil, err
			} else {
				return val, nil
			}
		}
	}
	return nil, errors.New("eval: no value")
}
