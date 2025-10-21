package lisp

import "testing"

func TestMatch(t *testing.T) {
	for _, in := range [][3]string{
		{"((X x))", "(quote X)", "'x"},
		{"((X (quote x)))", "(atom X)", "(atom 'x)"},
		{"((X x) (Y y))", "(eq X Y)", "(eq x y)"},
		{"((X x))", "(car X)", "(car x)"},
		{"((X x))", "(cdr X)", "(cdr x)"},
		{"((X x) (Y y))", "(cons X Y)", "(cons x y)"},
		{"((C ((x (quote nil)) ((quote t) (quote t)))))", "(cond . C)", "(cond (x 'nil) ('t 't))"},
		{"((F id) (P (x)) (E x))", "(label F (lambda P E))", "(label id (lambda (x) x))"},
		{"((P (x)) (E x) (A (quote t)))", "((lambda P E) A)", "((lambda (x) x) 't)"},
		{"((F id) (P (x)) (E x))", "(defun F P E)", "(defun id (x) x)"},
		{"((X x))", "X", "x"},
	} {
		pattern := Must(Read, in[1])
		expression := Must(Read, in[2])
		if expected, actual :=
			Must(Read, in[0]),
			Must2(Match, pattern, expression); expected.String() != actual.String() {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
