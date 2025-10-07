package lisp

import "strings"

func Read(source string) (Value, error) {
	return NewReader(NewTokenizer(strings.NewReader(source))).ReadExpression()
}
