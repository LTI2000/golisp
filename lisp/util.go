package lisp

import "strings"

func Read(source string) (Value, error) {
	return NewParser(NewTokenizer(strings.NewReader(source))).ReadExpression()
}
