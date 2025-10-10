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

func TestParseCarPrimitive(t *testing.T) {
	if expected, actual := "(car (quote (a b)))", Must(ParseExpression, Must(Read, "(car '(a b))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestParseCdrPrimitive(t *testing.T) {
	if expected, actual := "(cdr (quote (a b)))", Must(ParseExpression, Must(Read, "(cdr '(a b))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseConsPrimitive(t *testing.T) {
	if expected, actual := "(cons (quote a) (quote b))", Must(ParseExpression, Must(Read, "(cons 'a 'b)")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseCondPrimitive(t *testing.T) {
	if expected, actual := "(cond ((quote a) (quote b)) ((quote t) (quote u)))", Must(ParseExpression, Must(Read, "(cond ('a 'b) ('t 'u))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
