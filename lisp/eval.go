package lisp

import "fmt"

func Append(x, y Expression) (Expression, error) {
	if x == NIL {
		return y, nil
	} else if car_x, cdr_x, err := Uncons(x, "Append1"); err != nil {
		return nil, err
	} else if rest, err := Append(cdr_x, y); err != nil {
		return nil, err
	} else {
		return Cons(car_x, rest), nil
	}
}

func assoc(x, y Expression) (Expression, error) {
	if y == NIL {
		return nil, fmt.Errorf("%v: unbound variable", x)
	} else if car_y, cdr_y, err := Uncons(y, "assoc1"); err != nil {
		return nil, err
	} else if caar_y, cdar_y, err := Uncons(car_y, "assoc2"); err != nil {
		return nil, err
	} else if caar_y == x {
		if cadar_y, _, err := Uncons(cdar_y, "assoc3"); err != nil {
			return nil, err
		} else {
			return cadar_y, nil
		}
	} else {
		return assoc(x, cdr_y)
	}
}

func eval(exp Expression, env Environment) (Expression, error) {
	if ok := Match0("X:atom", exp); ok {
		return env.Lookup(exp)
	} else if x, ok := Match1("(quote X)", exp, "X"); ok {
		return x, nil
	} else if x, ok := Match1("(atom X)", exp, "X"); ok {
		if a0, err := eval(x, env); err != nil {
			return nil, err
		} else {
			return Bool(Atom(a0)), nil
		}
	} else if x, y, ok := Match2("(eq X Y)", exp, "X", "Y"); ok {
		if a0, err := eval(x, env); err != nil {
			return nil, err
		} else if b0, err := eval(y, env); err != nil {
			return nil, err
		} else {
			return Bool(Eq(a0, b0)), nil
		}
	} else if x, ok := Match1("(car X)", exp, "X"); ok {
		if a0, err := eval(x, env); err != nil {
			return nil, err
		} else {
			return Car(a0)
		}
	} else if x, ok := Match1("(cdr X)", exp, "X"); ok {
		if a0, err := eval(x, env); err != nil {
			return nil, err
		} else {
			return Cdr(a0)
		}
	} else if x, y, ok := Match2("(cons X Y)", exp, "X", "Y"); ok {
		if a0, err := eval(x, env); err != nil {
			return nil, err
		} else if b0, err := eval(y, env); err != nil {
			return nil, err
		} else {
			return Cons(a0, b0), nil
		}
	} else if c, ok := Match1("(cond . C:list)", exp, "C"); ok {
		return evcon(c, env)
	} else if x, y, ok := Match2("(X:atom . Y:list)", exp, "X", "Y"); ok {
		if f, err := env.Lookup(x); err != nil {
			return nil, err
		} else {
			return eval(Cons(f, y), env)
		}
	} else if n, f, p, ok := Match3("((label N F) . P:list)", exp, "N", "F", "P"); ok {
		return eval(Cons(f, p), Extend(n, Must(Car, exp), env))
	} else if p, b, x, ok := Match3("((lambda P B) . X:list)", exp, "P", "B", "X"); ok {
		if v0, err := evlis(x, env); err != nil {
			return nil, err
		} else {
			return eval(b, ExtendList(Slice(p), Slice(v0), env))
		}
	} else {
		return nil, fmt.Errorf("eval: malformed expression: %v", exp)
	}
}

func evcon(clauses Expression, env Environment) (Expression, error) {
	if test, consequent, alternates, ok := Match3("((TEST CONSEQUENT) . ALTERNATES:list)", clauses, "TEST", "CONSEQUENT", "ALTERNATES"); ok {
		if t, err := eval(test, env); err != nil {
			return nil, err
		} else if t != NIL {
			return eval(consequent, env)
		} else {
			return evcon(alternates, env)
		}
	} else {
		return nil, fmt.Errorf("evcon: malformed clauses: %v", clauses)
	}
}

func evlis(exps Expression, env Environment) (Expression, error) {
	if ok := Match0("()", exps); ok {
		return NIL, nil
	} else if head, tail, ok := Match2("(HEAD . TAIL:list)", exps, "HEAD", "TAIL"); ok {
		if first, err := eval(head, env); err != nil {
			return nil, err
		} else if rest, err := evlis(tail, env); err != nil {
			return nil, err
		} else {
			return Cons(first, rest), nil
		}
	} else {
		return nil, fmt.Errorf("evlis: malformed list: %v", exps)
	}
}

func Eval(exp Expression, env Environment) (Expression, error) {
	return eval(exp, env)
}
