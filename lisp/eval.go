package lisp

import "fmt"

func Append(x, y Expression) (Expression, error) {
	if x == NIL {
		return y, nil
	} else if car_x, cdr_x, err := Uncons(x); err != nil {
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
	} else if car_x, cdr_x, err := Uncons(x); err != nil {
		return nil, err
	} else if car_y, cdr_y, err := Uncons(y); err != nil {
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
	} else if car_y, cdr_y, err := Uncons(y); err != nil {
		return nil, err
	} else if caar_y, cdar_y, err := Uncons(car_y); err != nil {
		return nil, err
	} else if caar_y == x {
		if cadar_y, _, err := Uncons(cdar_y); err != nil {
			return nil, err
		} else {
			return cadar_y, nil
		}
	} else {
		return assoc(x, cdr_y)
	}
}

func Eval(e, a Expression) (Expression, error) {
	if Atom(e) {
		return assoc(e, a)
	} else if car_e, cdr_e, err := Uncons(e); err != nil {
		return nil, err
	} else if Atom(car_e) {
		if car_e == QUOTE {
			return Car(cdr_e)
		} else if car_e == ATOM {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, a); err != nil {
				return nil, err
			} else {
				return Bool(Atom(x)), nil
			}
		} else if car_e == EQ {
			if cadr_e, cddr_e, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if caddr_e, _, err := Uncons(cddr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, a); err != nil {
				return nil, err
			} else if y, err := Eval(caddr_e, a); err != nil {
				return nil, err
			} else {
				return Bool(Eq(x, y)), nil
			}
		} else if car_e == CAR {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, a); err != nil {
				return nil, err
			} else {
				return Car(x)
			}
		} else if car_e == CDR {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, a); err != nil {
				return nil, err
			} else {
				return Cdr(x)
			}
		} else if car_e == CONS {
			if cadr_e, cddr_e, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if caddr_e, _, err := Uncons(cddr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, a); err != nil {
				return nil, err
			} else if y, err := Eval(caddr_e, a); err != nil {
				return nil, err
			} else {
				return Cons(x, y), nil
			}
		} else if car_e == COND {
			return evcon(cdr_e, a)
		} else if f, err := assoc(car_e, a); err != nil {
			return nil, err
		} else {
			return Eval(Cons(f, cdr_e), a)
		}
	} else if caar_e, cdar_e, err := Uncons(car_e); err != nil {
		return nil, err
	} else if caar_e == LABEL {
		if cadar_e, cddar_e, err := Uncons(cdar_e); err != nil {
			return nil, err
		} else if caddar_e, _, err := Uncons(cddar_e); err != nil {
			return nil, err
		} else {
			return Eval(Cons(caddar_e, cdr_e), Cons(List(cadar_e, car_e), a))
		}
	} else if caar_e == LAMBDA {
		if cadar_e, cddar_e, err := Uncons(cdar_e); err != nil {
			return nil, err
		} else if caddar_e, _, err := Uncons(cddar_e); err != nil {
			return nil, err
		} else if args, err := evlis(cdr_e, a); err != nil {
			return nil, err
		} else if a1, err := pair(cadar_e, args); err != nil {
			return nil, err
		} else if a2, err := Append(a1, a); err != nil {
			return nil, err
		} else {
			return Eval(caddar_e, a2)
		}
	} else {
		return nil, fmt.Errorf("Eval: bad expression: %v", e)
	}
}

func evcon(c, a Expression) (Expression, error) {
	if car_c, cdr_c, err := Uncons(c); err != nil {
		return nil, err
	} else if caar_c, cdar_c, err := Uncons(car_c); err != nil {
		return nil, err
	} else if x, err := Eval(caar_c, a); err != nil {
		return nil, err
	} else if x != NIL {
		if cadar_c, _, err := Uncons(cdar_c); err != nil {
			return nil, err
		} else {
			return Eval(cadar_c, a)
		}
	} else {
		return evcon(cdr_c, a)
	}
}

func evlis(m, a Expression) (Expression, error) {
	if m == NIL {
		return NIL, nil
	} else if car_m, cdr_m, err := Uncons(m); err != nil {
		return nil, err
	} else if first, err := Eval(car_m, a); err != nil {
		return nil, err
	} else if rest, err := evlis(cdr_m, a); err != nil {
		return nil, err
	} else {
		return Cons(first, rest), nil
	}
}

func Defun(e Expression) (Expression, error) {
	if car_e, cdr_e, err := Uncons(e); err != nil {
		return nil, err
	} else if car_e == DEFUN {
		if cadr_e, cddr_e, err := Uncons(cdr_e); err != nil {
			return nil, err
		} else if caddr_e, cdddr_e, err := Uncons(cddr_e); err != nil {
			return nil, err
		} else if cadddr_e, _, err := Uncons(cdddr_e); err != nil {
			return nil, err
		} else {
			return List(cadr_e, List(LABEL, cadr_e, List(LAMBDA, caddr_e, cadddr_e))), nil
		}
	} else {
		return nil, fmt.Errorf("Defun: bad definition: %v", e)
	}
}
