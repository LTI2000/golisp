package lisp

import (
	"strings"
	"unicode"
)

var patterns map[string]Expression = make(map[string]Expression)

func makePattern(src string) Expression {
	pattern, ok := patterns[src]
	if !ok {
		pattern = Must(Read, src)
		patterns[src] = pattern
	}
	return pattern
}

func Match0(pattern string, expression Expression) bool {
	if _, ok := match(makePattern(pattern), expression); !ok {
		return false
	} else {
		return true
	}
}
func Match1(pattern string, expression Expression, name1 string) (Expression, bool) {
	if bindings, ok := match(makePattern(pattern), expression); !ok {
		return nil, false
	} else if value1, err := assoc(Symbol(name1), bindings); err != nil {
		return nil, false

	} else {
		return value1, true
	}
}

func Match2(pattern string, expression Expression, name1, name2 string) (Expression, Expression, bool) {
	if bindings, ok := match(makePattern(pattern), expression); !ok {
		return nil, nil, false
	} else if value1, err := assoc(Symbol(name1), bindings); err != nil {
		return nil, nil, false
	} else if value2, err := assoc(Symbol(name2), bindings); err != nil {
		return nil, nil, false
	} else {
		return value1, value2, true
	}
}

func Match3(pattern string, expression Expression, name1, name2, name3 string) (Expression, Expression, Expression, bool) {
	if bindings, ok := match(makePattern(pattern), expression); !ok {
		return nil, nil, nil, false
	} else if value1, err := assoc(Symbol(name1), bindings); err != nil {
		return nil, nil, nil, false
	} else if value2, err := assoc(Symbol(name2), bindings); err != nil {
		return nil, nil, nil, false
	} else if value3, err := assoc(Symbol(name3), bindings); err != nil {
		return nil, nil, nil, false
	} else {
		return value1, value2, value3, true
	}
}

func match(pattern, expression Expression) (Expression, bool) {
	switch p := pattern.(type) {
	case *symbol:
		return matchSymbolPattern(p, expression)
	case *cons:
		return matchPairPattern(p, expression)
	default:
		return nil, false
	}
}

func matchSymbolPattern(s *symbol, expression Expression) (Expression, bool) {
	if name, pred, ok := extractNameAndPredicate(s); ok {
		if pred(expression) {
			return List(List(name, expression)), true
		} else {
			return nil, false
		}
	} else if s == expression {
		return NIL, true
	} else {
		return nil, false
	}
}

func extractNameAndPredicate(s *symbol) (Expression, func(Expression) bool, bool) {
	if name, predicate, found := strings.Cut(s.name, ":"); found {
		if isUpperCase(name) {
			switch predicate {
			case "atom":
				return Symbol(name), Atom, true
			case "list":
				return Symbol(name), func(e Expression) bool { return e == NIL || !Atom(e) }, true
			default:
				panic("unknown match predicate: " + predicate)
			}
		} else {
			panic("bad match symbol: " + s.String())
		}
	} else {
		return s, func(Expression) bool { return true }, isUpperCase(s.name)
	}
}

func matchPairPattern(c *cons, expression Expression) (Expression, bool) {
	switch e := expression.(type) {
	case *cons:
		if b1, ok := match(c.car, e.car); !ok {
			return nil, false
		} else if b2, ok := match(c.cdr, e.cdr); !ok {
			return nil, false
		} else if b, err := Append(b1, b2); err != nil {
			return nil, false
		} else {
			return b, true
		}
	default:
		return nil, false
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
