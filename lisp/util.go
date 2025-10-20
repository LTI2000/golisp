package lisp

import (
	"strings"

	"github.com/LTI2000/golisp/lisp/scan"
)

func Read(source string) (Expression, error) {
	return NewReader(scan.NewScanner(strings.NewReader(source))).ReadValue()
}

func Must[A, B any](f func(A) (B, error), a A) B {
	if b, err := f(a); err != nil {
		panic("Must() failed: " + err.Error())
	} else {
		return b
	}
}

func Compose[A1, B1, A2, B2 any](f func(A1) (B1, error), g func(A2) (B2, error)) func(A1, A2) (B1, B2, error) {
	var b1 B1
	var b2 B2
	var err error = nil
	return func(a1 A1, a2 A2) (B1, B2, error) {
		if err == nil {
			b1, err = f(a1)
		}
		if err == nil {
			b2, err = g(a2)
		}
		return b1, b2, err
	}
}
