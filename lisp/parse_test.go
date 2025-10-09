package lisp

import "testing"

func TestParseQuotePrimitive(t *testing.T) {
	if expected, actual := "(quote a)", Must(ParseExpression, Must(Read, "'a")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestParseAtomPrimitive(t *testing.T) {
	if expected, actual := "(atom (quote a))", Must(ParseExpression, Must(Read, "(atom 'a)")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseEqPrimitive(t *testing.T) {
	if expected, actual := "(eq (quote a) (quote b))", Must(ParseExpression, Must(Read, "(eq 'a 'b)")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
