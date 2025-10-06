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
	value := Cons(nil, nil)
	actual := value.IsAtom()
	if actual {
		t.Errorf("%v", value)
	}
}

func TestConsCar(t *testing.T) {
	car := Symbol("car")
	value := Cons(car, nil)
	actual := value.GetCar()
	if actual != car {
		t.Errorf("%v", value)
	}
}

func TestConsCdr(t *testing.T) {
	cdr := Symbol("cdr")
	value := Cons(nil, cdr)
	actual := value.GetCdr()
	if actual != cdr {
		t.Errorf("%v", value)
	}
}

func TestConsString(t *testing.T) {
	if expected, actual := "(x)", Cons(Symbol("x"), Nil).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(x y)", Cons(Symbol("x"), Cons(Symbol("y"), Nil)).String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
	if expected, actual := "(x . y)", Cons(Symbol("x"), Symbol("y")).String(); expected != actual {
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
