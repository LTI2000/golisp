package lisp

import (
	"testing"
)

func TestReadSymbol(t *testing.T) {
	if value, err := Read("quote"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("quote"), value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadQuotation(t *testing.T) {
	if value, err := Read("'a"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Quote, value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("a"), value.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, value.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadEmptyList(t *testing.T) {
	if value, err := Read("()"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Nil, value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadSingletonList(t *testing.T) {
	if value, err := Read("(x)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, value.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadProperList(t *testing.T) {
	if value, err := Read("(x y)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), value.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, value.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadNestedList(t *testing.T) {
	if value, err := Read("(x (y) z)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), value.GetCdr().GetCar().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, value.GetCdr().GetCar().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("z"), value.GetCdr().GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, value.GetCdr().GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadDottedList(t *testing.T) {
	if value, err := Read("(x . y)"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), value.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadDottedListPair(t *testing.T) {
	if value, err := Read("((x . y) (a . b))"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Must(Read, "(x . y)").String(), value.GetCar().String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Must(Read, "((a . b))").String(), value.GetCdr().String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadDottedProperList(t *testing.T) {
	if value, err := Read("(x . (y . ())"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), value.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), value.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestReadLongDottedProperList(t *testing.T) {
	if value, err := Read("(x . (y . (z . ()))"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := "(x y z)", value.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestReadError(t *testing.T) {
	if _, err := Read(")"); err == nil {
		t.Fatalf("expetced error")
	} else if expected, actual := "illegal token: ')'", err.Error(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func FuzzReadvalue(f *testing.F) {
	for _, seed := range []string{
		"x", "(x)", "(x y)", "(x . y)", "(x y ...)", "()",
		"(QUOTE x)", "(ATOM x)", "(EQ x y)", "(CAR x)", "(CDR x)", "(CONS x y)", "(COND (p e) ...)"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in string) {
		if value, err := Read(in); err != nil {
			t.Fatalf("err %v", err)
		} else if expected, actual := in, value.String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	})
}
