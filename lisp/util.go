package lisp

import "strings"

func Read(source string) (Value, error) {
	return NewReader(NewScanner(strings.NewReader(source))).ReadValue()
}
