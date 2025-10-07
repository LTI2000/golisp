package lisp

import "testing"

func TestMatch(t *testing.T) {
	if pattern, err := Read("()"); err != nil {
		t.Fatalf("err %v", err)
	} else if value, err := Read("nil"); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := true, match(pattern, value); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
