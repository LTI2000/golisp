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
	} {
		if expected, actual :=
			Must(Read, in[0]),
			Must2(Eval, Must(Read, in[1]), Must(Read, in[2])); expected.String() != actual.String() {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
