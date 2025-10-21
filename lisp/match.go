package lisp

import (
	"errors"
	"unicode"
)

func Match(pattern, expression Expression) (Expression, error) {
	return match(pattern, expression)
}

func match(pattern, expression Expression) (Expression, error) {
	switch p := pattern.(type) {
	case *symbol:
		return matchSymbolPattern(p, expression)
	case *cons:
		return matchPairPattern(p, expression)
	}
	panic("match: unknown type")
}

func matchSymbolPattern(s *symbol, expression Expression) (Expression, error) {
	if isUpperCase(s.name) {
		return List(List(s, expression)), nil
	} else if s == expression {
		return NIL, nil
	} else {
		return nil, errors.New("nomatch")
	}
}

func matchPairPattern(c *cons, expression Expression) (Expression, error) {
	switch e := expression.(type) {
	case *cons:
		if b1, err := match(c.car, e.car); err != nil {
			return nil, err
		} else if b2, err := match(c.cdr, e.cdr); err != nil {
			return nil, err
		} else if b, err := Append(b1, b2); err != nil {
			return nil, err
		} else {
			return b, nil
		}
	default:
		return nil, errors.New("nomatch")
	}
}

// utilities
func isUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
