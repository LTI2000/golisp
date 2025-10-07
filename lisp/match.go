package lisp

func match(pattern Value, value Value) bool {
	return pattern.IsEq(value)
}
