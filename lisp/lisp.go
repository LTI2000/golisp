package lisp

func QUOTE(x Value) Value {
	return x
}

func ATOM(x Value) Value {
	return BoolSymbol(x.IsAtom())
}

func EQ(x Value, y Value) Value {
	return BoolSymbol(x.IsEq(y))
}

func CAR(x Value) Value {
	return x.GetCar()
}

func CDR(x Value) Value {
	return x.GetCdr()
}

func CONS(x Value, y Value) Value {
	return Cons(x, y)
}

func COND(cs []Value) Value {
	panic("unimplemented")
}
