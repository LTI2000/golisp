package lisp

import "testing"

func TestEval(t *testing.T) {
	env := NewEnvironment()
	for _, in := range []struct {
		res string
		exp string
		env Environment
	}{
		{"a", "x", Extend("x", Symbol("a"), env)},
		{"a", "(quote a)", env},
		{"t", "(atom 'a)", env},
		{"nil", "(atom '(a b))", env},
		{"t", "(eq 'a 'a)", env},
		{"nil", "(eq 'a 'b)", env},
		{"a", "(car '(a . b))", env},
		{"b", "(cdr '(a . b))", env},
		{"(a . b)", "(cons 'a 'b)", env},
		{"no", "(cond ((atom '(a b)) 'yes) ('t 'no))", env},
		{"(a . b)", "(kons 'a 'b)", Extend("kons", Symbol("cons"), env)},
		{
			`(a m (a m c) d)`,
			`((label subst 
                     (lambda (x y z)
                       (cond ((atom z)
                              (cond ((eq z y) x)
                                    ('t z)))
                             ('t (cons (subst x y (car z))
                                       (subst x y (cdr z))))))) 'm 'b '(a b (a b c) d))`,
			env,
		},
	} {
		if expected, actual :=
			Must(Read, in.res),
			Must2(Eval, Must(Read, in.exp), in.env); expected.String() != actual.String() {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
