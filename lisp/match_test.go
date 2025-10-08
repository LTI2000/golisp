package lisp

import "testing"

func TestMatchSymbol(t *testing.T) {
	pattern := Must(Read, "()")
	value := Must(Read, "nil")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchPair(t *testing.T) {
	if pattern, err := Read("(a b)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("(a . (b . ()))"); err != nil {
		t.Fatalf("err %v", err)
	} else if bindings, matches := NewPattern(pattern).Match(NewBindings(), value); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariable(t *testing.T) {
	if pattern, err := Read("X"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("t"); err != nil {
		t.Fatalf("err %v", err)
	} else if bindings, matches := NewPattern(pattern).Match(NewBindings(), value); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[X=t]}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair1(t *testing.T) {
	if pattern, err := Read("(X)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("t"); err != nil {
		t.Fatalf("err %v", err)
	} else if bindings, matches := NewPattern(pattern).Match(NewBindings(), value); matches {
		t.Errorf("expected non-match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair2(t *testing.T) {
	if pattern, err := Read("(X Y)"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("(y x)"); err != nil {
		t.Fatalf("err %v", err)
	} else if bindings, matches := NewPattern(pattern).Match(NewBindings(), value); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[Y=x][X=y]}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
