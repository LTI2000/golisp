package lisp

func match(pattern Value, expression Value) bool {
	return pattern.IsEq(expression)
}
