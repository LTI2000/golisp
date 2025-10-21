package lisp

func Must[A, R any](f func(A) (R, error), a A) R {
	if r, err := f(a); err != nil {
		panic("Must() failed: " + err.Error())
	} else {
		return r
	}
}

func Must2[A, B, R any](f func(A, B) (R, error), a A, b B) R {
	if r, err := f(a, b); err != nil {
		panic("Must() failed: " + err.Error())
	} else {
		return r
	}
}

func Compose[A1, B1, A2, B2 any](f func(A1) (B1, error), g func(A2) (B2, error)) func(A1, A2) (B1, B2, error) {
	var b1 B1
	var b2 B2
	var err error = nil
	return func(a1 A1, a2 A2) (B1, B2, error) {
		if err == nil {
			b1, err = f(a1)
		}
		if err == nil {
			b2, err = g(a2)
		}
		return b1, b2, err
	}
}
