package lisp

import "testing"

func TestEval(t *testing.T) {
	expected := Must(Read, "a")
	if actual, err := Eval(Must(Read, "x"), Must(Read, "((x a))")); err != nil {
		t.Fatalf("err %v", err)

	} else if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)

	}
}
