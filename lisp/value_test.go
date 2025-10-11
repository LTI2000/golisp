package lisp

import (
	"testing"
)

func TestSymbolIdentity(t *testing.T) {
	if Symbol("a") != Symbol("a") {
		t.Errorf("same symbols are not identical")
	}
	if Symbol("a") == Symbol("b") {
		t.Errorf("different symbols are identical")
	}
}

func TestSymbolAtom(t *testing.T) {
	value := Symbol("sym")
	actual := value.IsAtom()
	if !actual {
		t.Errorf("%v", value)
	}
}

func TestSymbolCar(t *testing.T) {
	var recovered = false
	defer func() {
		if r := recover(); r != nil {
			recovered = true
		}
	}()
	sym := Symbol("sym")
	sym.GetCar()
	if !recovered {
		t.Errorf("%v", sym)
	}
}

func TestSymbolCdr(t *testing.T) {
	var recovered = false
	defer func() {
		if r := recover(); r != nil {
			recovered = true
		}
	}()
	sym := Symbol("sym")
	sym.GetCdr()
	if !recovered {
		t.Errorf("%v", sym)
	}
}

func TestSymbolString(t *testing.T) {
	if expected, actual := "x", Symbol("x").String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "()", Nil.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestConsAtom(t *testing.T) {
	value := Pair(nil, nil)
	actual := value.IsAtom()
	if actual {
		t.Errorf("%v", value)
	}
}

func TestConsCar(t *testing.T) {
	car := Symbol("car")
	value := Pair(car, nil)
	actual := value.GetCar()
	if actual != car {
		t.Errorf("%v", value)
	}
}

func TestConsCdr(t *testing.T) {
	cdr := Symbol("cdr")
	value := Pair(nil, cdr)
	actual := value.GetCdr()
	if actual != cdr {
		t.Errorf("%v", value)
	}
}

func TestConsString(t *testing.T) {
	if expected, actual := "(x)", Pair(Symbol("x"), Nil).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(x y)", Pair(Symbol("x"), Pair(Symbol("y"), Nil)).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(x . y)", Pair(Symbol("x"), Symbol("y")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestList(t *testing.T) {
	if expected, actual := "()", List().String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a)", List(Symbol("a")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a b c)", List(Symbol("a"), Symbol("b"), Symbol("c")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestSlice(t *testing.T) {
	var slice []Value

	slice = Slice(List())
	if expected, actual := 0, len(slice); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	slice = Slice(List(T))
	if expected, actual := 1, len(slice); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := T, slice[0]; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	slice = Slice(List(T, Nil, Quote))
	if expected, actual := 3, len(slice); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := T, slice[0]; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := Nil, slice[1]; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := Quote, slice[2]; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestConcat(t *testing.T) {
	if expected, actual := "()", Concat(Must(Read, "()")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a)", Concat(Must(Read, "((a))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a)", Concat(Must(Read, "(() (a))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a b c)", Concat(Must(Read, "(() (a) (b c))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(a b c)", Concat(Must(Read, "(() (a) () (b) () (c))")).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
