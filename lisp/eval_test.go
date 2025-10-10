package lisp

import "testing"

func TestEvalQuotePrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(quote a)")))
		if expected, actual := "a", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
func TestEvalAtomPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(atom 'a)")))
		if expected, actual := "t", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(atom '(a b))")))
		if expected, actual := "()", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}

func TestEvalEqPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(eq 'a 'a)")))
		if expected, actual := "t", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(eq 'a 'b)")))
		if expected, actual := "()", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(eq 'a '(b c))")))
		if expected, actual := "()", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(eq '(a b) 'c)")))
		if expected, actual := "()", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
func TestEvalCarPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(car '(a . b))")))
		if expected, actual := "a", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}

func TestEvalCdrPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(cdr '(a . b))")))
		if expected, actual := "b", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}

func TestEvalConsPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(cons 'a 'b)")))
		if expected, actual := "(a . b)", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}

func TestEvalCondPrimitive(t *testing.T) {
	{
		value := Must(Eval, Must(ParseExpression, Must(Read, "(cond ((atom '(a b)) 'no) ('t 'yes))")))
		if expected, actual := "yes", value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
