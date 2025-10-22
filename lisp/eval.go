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

func pair(x, y Expression) (Expression, error) {
	if x == NIL && y == NIL {
		return NIL, nil
	} else if car_x, cdr_x, err := Uncons(x, "pair1"); err != nil {
		return nil, err
	} else if car_y, cdr_y, err := Uncons(y, "pair2"); err != nil {
		return nil, err
	} else if rest, err := pair(cdr_x, cdr_y); err != nil {
		return nil, err
	} else {
		return Cons(List(car_x, car_y), rest), nil
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

func eval(e, a Expression) (Expression, error) {
	if ok := Match0("X:atom", e); ok {
		return assoc(e, a)
	} else if x, ok := Match1("(quote X)", e, "X"); ok {
		return x, nil
	} else if x, ok := Match1("(atom X)", e, "X"); ok {
		if a0, err := eval(x, a); err != nil {
			return nil, err
		} else {
			return Bool(Atom(a0)), nil
		}
	} else if x, y, ok := Match2("(eq X Y)", e, "X", "Y"); ok {
		if a0, err := eval(x, a); err != nil {
			return nil, err
		} else if b0, err := eval(y, a); err != nil {
			return nil, err
		} else {
			return Bool(Eq(a0, b0)), nil
		}
	} else if x, ok := Match1("(car X)", e, "X"); ok {
		if a0, err := eval(x, a); err != nil {
			return nil, err
		} else {
			return Car(a0)
		}
	} else if x, ok := Match1("(cdr X)", e, "X"); ok {
		if a0, err := eval(x, a); err != nil {
			return nil, err
		} else {
			return Cdr(a0)
		}
	} else if x, y, ok := Match2("(cons X Y)", e, "X", "Y"); ok {
		if a0, err := eval(x, a); err != nil {
			return nil, err
		} else if b0, err := eval(y, a); err != nil {
			return nil, err
		} else {
			return Cons(a0, b0), nil
		}
	} else if c, ok := Match1("(cond . C:list)", e, "C"); ok {
		return evcon(c, a)
	} else if x, y, ok := Match2("(X:atom . Y:list)", e, "X", "Y"); ok {
		if f, err := assoc(x, a); err != nil {
			return nil, err
		} else {
			return eval(Cons(f, y), a)
		}
	} else if n, f, p, ok := Match3("((label N F) . P:list)", e, "N", "F", "P"); ok {
		return eval(Cons(f, p), Cons(List(n, Must(Car, e)), a))
	} else if p, b, x, ok := Match3("((lambda P B) . X:list)", e, "P", "B", "X"); ok {
		if v0, err := evlis(x, a); err != nil {
			return nil, err
		} else if v1, err := pair(p, v0); err != nil {
			return nil, err
		} else if v2, err := Append(v1, a); err != nil {
			return nil, err
		} else {
			return eval(b, v2)
		}
	} else {
		return nil, fmt.Errorf("eval: illegal expressoin: %v", e)
	}
}

func evcon(c, a Expression) (Expression, error) {
	if car_c, cdr_c, err := Uncons(c, "evcon1"); err != nil {
		return nil, err
	} else if caar_c, cdar_c, err := Uncons(car_c, "evcon2"); err != nil {
		return nil, err
	} else if x, err := eval(caar_c, a); err != nil {
		return nil, err
	} else if x != NIL {
		if cadar_c, _, err := Uncons(cdar_c, "evcon3"); err != nil {
			return nil, err
		} else {
			return eval(cadar_c, a)
		}
	} else {
		return evcon(cdr_c, a)
	}
}

func evlis(m, a Expression) (Expression, error) {
	if m == NIL {
		return NIL, nil
	} else if car_m, cdr_m, err := Uncons(m, "evlis1"); err != nil {
		return nil, err
	} else if first, err := eval(car_m, a); err != nil {
		return nil, err
	} else if rest, err := evlis(cdr_m, a); err != nil {
		return nil, err
	} else {
		return Cons(first, rest), nil
	}
}

func Eval(e, a Expression) (Expression, error) {
	return eval(e, a)
}
