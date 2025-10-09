package lisp

import "testing"

func TestEvalLiteralExpression(t *testing.T) {
	value := Must(Eval, Must(ParseExpression, Must(Read, "(quote a)")))
	if expected, actual := "a", value.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
