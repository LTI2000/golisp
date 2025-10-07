package lisp

import (
	"testing"
)

func TestReadSymbol(t *testing.T) {
	if expression, err := Read("quote"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("quote"), expression; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadQuotation(t *testing.T) {
	if expression, err := Read("'a"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Quote, expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("a"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadEmptyList(t *testing.T) {
	if expression, err := Read("()"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Nil, expression; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadSingletonList(t *testing.T) {
	if expression, err := Read("(x)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadProperList(t *testing.T) {
	if expression, err := Read("(x y)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadNestedList(t *testing.T) {
	if expression, err := Read("(x (y) z)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCar().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("z"), expression.GetCdr().GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadDottedList(t *testing.T) {
	if expression, err := Read("(x . y)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadDottedProperList(t *testing.T) {
	if expression, err := Read("(x . (y . ())"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestReadLongDottedProperList(t *testing.T) {
	if expression, err := Read("(x . (y . (z . ()))"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := "(x y z)", expression.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func FuzzReadExpression(f *testing.F) {
	for _, seed := range []string{
		"x", "(x)", "(x y)", "(x . y)", "(x y ...)", "()",
		"(QUOTE x)", "(ATOM x)", "(EQ x y)", "(CAR x)", "(CDR x)", "(CONS x y)", "(COND ((p e) ...))"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in string) {
		if expression, err := Read(in); err != nil {
			t.Fatalf("err %v", err)
		} else if expected, actual := in, expression.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}
