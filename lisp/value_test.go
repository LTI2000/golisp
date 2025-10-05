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
	actual := value.Atom()
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
	sym.Car()
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
	sym.Cdr()
	if !recovered {
		t.Errorf("%v", sym)
	}
}
func TestConsAtom(t *testing.T) {
	value := Cons(nil, nil)
	actual := value.Atom()
	if actual {
		t.Errorf("%v", value)
	}
}

func TestConsCar(t *testing.T) {
	car := Symbol("car")
	value := Cons(car, nil)
	actual := value.Car()
	if actual != car {
		t.Errorf("%v", value)
	}
}
func TestConsCdr(t *testing.T) {
	cdr := Symbol("cdr")
	value := Cons(nil, cdr)
	actual := value.Cdr()
	if actual != cdr {
		t.Errorf("%v", value)
	}
}
