package lisp

import "testing"

func TestEval(t *testing.T) {
	for _, in := range [][3]string{
		{"a", "x", "((x a))"},
		{"a", "(quote a)", "()"},
		{"t", "(atom 'a)", "()"},
		{"()", "(atom '(a b))", "()"},
	} {
		if expected, actual :=
			Must(Read, in[0]),
			Must2(Eval, Must(Read, in[1]), Must(Read, in[2])); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
