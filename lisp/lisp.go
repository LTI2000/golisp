package lisp

func Quote(x Value) Value {
	return x
}

func Atom(x Value) Value {
	if x.Atom() {
		return Symbol("t")
	} else {
		return Symbol("nil")
	}
}

func Eq(x Value, y Value) Value {
	if x.Eq(y) {
		return Symbol("t")
	} else {
		return Symbol("nil")
	}
}

func Car(x Value) Value {
	return x.Car()
}

func Cdr(x Value) Value {
	return x.Cdr()
}

func Cons(x Value, y Value) Value {
	return Pair(x, y)
}

func Cond(cs []Value) Value {
	for _, c := range cs {
		return Atom(c)
	}
	return Symbol("nil")
}
