package lisp

import "testing"

func TestParseQuote(t *testing.T) {
	value := Must(Read, "'a")

	if expression, err := ParseExpression(value); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := "(quote a)", expression.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
