package lisp

import (
	"testing"
)

// test expressions that write the same as they read
func TestRead1(t *testing.T) {
	for _, in := range []string{
		"quote", "()", "(x)", "(x y)", "(x (y) z)", "(x . y)", "((x . y) (a . b))", "...",
	} {
		if expected, actual := in, Must(Read, in).String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}

// test expressions that write different than they read, such as 'x, or (x . ())
func TestRead2(t *testing.T) {
	for _, in := range [][2]string{
		{"'a",
			"(quote a)"},
		{"(x . ())",
			"(x)"},
		{"(x . (y . ()))",
			"(x y)"},
		{"(x . (y . (z . ())))",
			"(x y z)"},
	} {
		if expected, actual := in[1], Must(Read, in[0]).String(); expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
func TestReadError(t *testing.T) {
	if _, err := Read(")"); err == nil {
		t.Fatalf("expected error")
	} else if expected, actual := "illegal token: ')'", err.Error(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
