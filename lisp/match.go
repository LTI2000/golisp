package lisp

import (
	"strings"
	"unicode"
)

var patternCache map[string]Expression = make(map[string]Expression)

func makePattern(src string) Expression {
	pattern, ok := patternCache[src]
	if !ok {
		pattern = Must(Read, src)
		patternCache[src] = pattern
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
	} else if value1, err := bindings.Lookup(name1); err != nil {
		return nil, false
	} else {
		return value1, true
	}
}

func Match2(pattern string, expression Expression, name1, name2 string) (Expression, Expression, bool) {
	if bindings, ok := match(makePattern(pattern), expression); !ok {
		return nil, nil, false
	} else if value1, err := bindings.Lookup(name1); err != nil {
		return nil, nil, false
	} else if value2, err := bindings.Lookup(name2); err != nil {
		return nil, nil, false
	} else {
		return value1, value2, true
	}
}

func Match3(pattern string, expression Expression, name1, name2, name3 string) (Expression, Expression, Expression, bool) {
	if bindings, ok := match(makePattern(pattern), expression); !ok {
		return nil, nil, nil, false
	} else if value1, err := bindings.Lookup(name1); err != nil {
		return nil, nil, nil, false
	} else if value2, err := bindings.Lookup(name2); err != nil {
		return nil, nil, nil, false
	} else if value3, err := bindings.Lookup(name3); err != nil {
		return nil, nil, nil, false
	} else {
		return value1, value2, value3, true
	}
}

func match(pattern, expression Expression) (Environment, bool) {
	switch p := pattern.(type) {
	case *symbol:
		return matchSymbolPattern(p, expression)
	case *cons:
		return matchPairPattern(p, expression)
	default:
		return nil, false
	}
}

func matchSymbolPattern(s *symbol, expression Expression) (Environment, bool) {
	if name, pred, ok := extractNameAndPredicate(s); ok {
		if pred(expression) {
			return Extend(name, expression, NewEnvironment()), true
		} else {
			return nil, false
		}
	} else if s == expression {
		return NewEnvironment(), true
	} else {
		return nil, false
	}
}

func matchPairPattern(c *cons, expression Expression) (Environment, bool) {
	switch e := expression.(type) {
	case *cons:
		if b1, ok := match(c.car, e.car); !ok {
			return nil, false
		} else if b2, ok := match(c.cdr, e.cdr); !ok {
			return nil, false
		} else {
			return Merge(b1, b2), true
		}
	default:
		return nil, false
	}
}

// utilities

func extractNameAndPredicate(s *symbol) (string, func(Expression) bool, bool) {
	if name, predicate, found := strings.Cut(s.name, ":"); found {
		if isUpperCase(name) {
			switch predicate {
			case "atom":
				return name, Atom, true
			case "list":
				return name, func(e Expression) bool { return e == NIL || !Atom(e) }, true
			default:
				panic("unknown match predicate: " + predicate)
			}
		} else {
			panic("bad match symbol: " + s.String())
		}
	} else {
		return s.String(), func(Expression) bool { return true }, isUpperCase(s.name)
	}
}

func isUpperCase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
