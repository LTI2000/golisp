package lisp

import "fmt"

// (defun null. (x)
//   (eq x '()))

// (defun and. (x y)
//   (cond (x (cond (y 't) ('t '())))
//         ('t '())))

// (defun not. (x)
//   (cond (x '())
//         ('t 't)))

// (defun append. (x y)
//   (cond ((null. x) y)
//         ('t (cons (car x) (append. (cdr x) y)))))

// (defun pair. (x y)
//   (cond ((and. (null. x) (null. y)) '())
//         ((and. (not. (atom x)) (not. (atom y)))
//          (cons (list (car x) (car y))
//                (pair. (cdr x) (cdr y))))))

// (defun assoc. (x y)
//   (cond ((eq (caar y) x) (cadar y))
//         ('t (assoc. x (cdr y)))))

func assoc(x, y Expression) (Expression, error) {
	if car_y, cdr_y, err := Uncons(y); err != nil {
		return nil, fmt.Errorf("%v: unbound variable", x)
	} else if caar_y, cdar_y, err := Uncons(car_y); err != nil {
		return nil, err
	} else if Eq(caar_y, x) {
		return Car(cdar_y)
	} else {
		return assoc(x, cdr_y)
	}
}

// (defun eval. (e a)
//
//	(cond
//	  ((atom e) (assoc. e a))
//	  ((atom (car e))
//	   (cond
//	     ((eq (car e) 'quote) (cadr e))
//	     ((eq (car e) 'atom)  (atom (eval. (cadr e) a)))
//	     ((eq (car e) 'eq)    (eq   (eval. (cadr e) a)
//	                                (eval. (caddr e) a)))
//	     ((eq (car e) 'car)   (car  (eval. (cadr e) a)))
//	     ((eq (car e) 'cdr)   (cdr  (eval. (cadr e) a)))
//	     ((eq (car e) 'cons)  (cons (eval. (cadr e) a)
//	                                (eval. (caddr e) a)))
//	     ((eq (car e) 'cond)  (evcon. (cdr e) a))
//	     ('t (eval. (cons (assoc. (car e) a)
//	                      (cdr e))
//	                a))))
//	  ((eq (caar e) 'label)
//	   (eval. (cons (caddar e) (cdr e))
//	          (cons (list (cadar e) (car e)) a)))
//	  ((eq (caar e) 'lambda)
//	   (eval. (caddar e)
//	          (append. (pair. (cadar e) (evlis. (cdr e) a))
//	                   a)))))

func Eval(e, a Expression) (Expression, error) {
	if Atom(e) {
		return assoc(e, a)
	} else if car_e, cdr_e, err := Uncons(e); err != nil {
		return nil, err
	} else if Atom(car_e) {
		if Eq(car_e, Symbol("quote")) {
			return Car(cdr_e)
		} else if Eq(car_e, Symbol("atom")) {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, e); err != nil {
				return nil, err
			} else {
				return Bool(Atom(x)), nil
			}
		} else if Eq(car_e, Symbol("eq")) {
			if cadr_e, cddr_e, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if caddr_e, _, err := Uncons(cddr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, e); err != nil {
				return nil, err
			} else if y, err := Eval(caddr_e, e); err != nil {
				return nil, err
			} else {
				return Bool(Eq(x, y)), nil
			}
		} else if Eq(car_e, Symbol("car")) {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, e); err != nil {
				return nil, err
			} else {
				return Car(x)
			}
		} else if Eq(car_e, Symbol("cdr")) {
			if cadr_e, _, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, e); err != nil {
				return nil, err
			} else {
				return Cdr(x)
			}
		} else if Eq(car_e, Symbol("cons")) {
			if cadr_e, cddr_e, err := Uncons(cdr_e); err != nil {
				return nil, err
			} else if caddr_e, _, err := Uncons(cddr_e); err != nil {
				return nil, err
			} else if x, err := Eval(cadr_e, e); err != nil {
				return nil, err
			} else if y, err := Eval(caddr_e, e); err != nil {
				return nil, err
			} else {
				return Cons(x, y), nil
			}
		} else if Eq(car_e, Symbol("cond")) {
			return evcon(cdr_e, a)
		} else if f, err := assoc(car_e, a); err != nil {
			return nil, err
		} else {
			return Eval(Cons(f, cdr_e), a)
		}
	} else {
		return nil, fmt.Errorf("Eval: bad expression: %v", e)
	}
}

// (defun evcon. (c a)
//   (cond ((eval. (caar c) a)
//          (eval. (cadar c) a))
//         ('t (evcon. (cdr c) a))))

func evcon(c, a Expression) (Expression, error) {
	if car_c, cdr_c, err := Uncons(c); err != nil {
		return nil, err
	} else if caar_c, cdar_c, err := Uncons(car_c); err != nil {
		return nil, err
	} else if x, err := Eval(caar_c, a); err != nil {
		return nil, err
	} else if x != Nil {
		if cadar_c, _, err := Uncons(cdar_c); err != nil {
			return nil, err
		} else {
			return Eval(cadar_c, a)
		}
	} else {
		return evcon(cdr_c, a)
	}
}

// (defun evlis. (m a)
//   (cond ((null. m) '())
//         ('t (cons (eval.  (car m) a)
//                   (evlis. (cdr m) a)))))
