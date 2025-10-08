package lisp

import "testing"

func TestMatchSymbol(t *testing.T) {
	if pattern, err := Read("()"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("nil"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, NewPattern(pattern).Match(value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestMatchPair(t *testing.T) {
	if pattern, err := Read("(a b)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("(a . (b . ()))"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, NewPattern(pattern).Match(value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestMatchVariable(t *testing.T) {
	if pattern, err := Read("X"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("t"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, NewPattern(pattern).Match(value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair1(t *testing.T) {
	if pattern, err := Read("(X)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("t"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, !NewPattern(pattern).Match(value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestMatchVariablePair2(t *testing.T) {
	if pattern, err := Read("(X)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("(t)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, NewPattern(pattern).Match(value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
