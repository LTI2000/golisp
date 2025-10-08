package lisp

import "strings"

func Read(source string) (Value, error) {
	return NewReader(NewScanner(strings.NewReader(source))).ReadValue()
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
