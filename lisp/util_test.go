package lisp

import (
	"errors"
	"testing"
)

var err0 = errors.New("e0")
var err1 = errors.New("e1")
var f0 = func(x float64) (float64, error) {
	return x * 2.0, nil
}
var f1 = func(x float64) (float64, error) {
	return 0.0, err0
}
var g0 = func(s string) (int, error) {
	return len(s), nil
}
var g1 = func(s string) (int, error) {
	return 0, err1
}

func TestCompose(t *testing.T) {
	{
		x, l, err := Compose(f0, g0)(1.0, "23")
		if expected, actual := 2.0, x; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
		if expected, actual := 2, l; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
		if expected, actual := error(nil), err; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		_, _, err := Compose(f1, g0)(1.0, "23")
		if expected, actual := err0, err; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		_, _, err := Compose(f0, g1)(1.0, "23")
		if expected, actual := err1, err; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
	{
		_, _, err := Compose(f1, g1)(1.0, "23")
		if expected, actual := err0, err; expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
