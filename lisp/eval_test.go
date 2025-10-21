package lisp

import "testing"

func TestEval(t *testing.T) {
	for _, in := range [][3]string{
		{"a", "x", "((x a))"},
		{"a", "(quote a)", "()"},
		{"t", "(atom 'a)", "()"},
		{"nil", "(atom '(a b))", "()"},
		{"t", "(eq 'a 'a)", "()"},
		{"nil", "(eq 'a 'b)", "()"},
		{"a", "(car '(a . b))", "()"},
		{"b", "(cdr '(a . b))", "()"},
		{"(a . b)", "(cons 'a 'b)", "()"},
		{"no", "(cond ((atom '(a b)) 'yes) ('t 'no))", "()"},
		{"(a . b)", "(kons 'a 'b)", "((kons cons))"},
		{
			`(a m (a m c) d)`,
			`((label subst 
                     (lambda (x y z)
                       (cond ((atom z)
                              (cond ((eq z y) x)
                                    ('t z)))
                             ('t (cons (subst x y (car z))
                                       (subst x y (cdr z))))))) 'm 'b '(a b (a b c) d))`,
			`()`,
		},
	} {
		if expected, actual :=
			Must(Read, in[0]),
			Must2(Eval, Must(Read, in[1]), Must(Read, in[2])); expected.String() != actual.String() {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
