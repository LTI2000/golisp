package lisp

import "testing"

func TestMatchSymbol(t *testing.T) {
	pattern := Must(Read, "()")
	value := Must(Read, "nil")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchPair(t *testing.T) {
	pattern := Must(Read, "(a b)")
	value := Must(Read, "(a . (b . ()))")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariable(t *testing.T) {
	pattern := Must(Read, "X")
	value := Must(Read, "t")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[X=t]}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair1(t *testing.T) {
	pattern := Must(Read, "(X)")
	value := Must(Read, "t")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); matches {
		t.Errorf("expected non-match: %v %v", pattern, value)
	} else if expected, actual := "{}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair2(t *testing.T) {
	pattern := Must(Read, "(X Y)")
	value := Must(Read, "(y x)")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[Y=x][X=y]}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchCond(t *testing.T) {
	pattern := Must(Read, "(cond (P E) ...)")
	value := Must(Read, "(cond ('() 'no) ('t 'yes))")

	if bindings, matches := NewPattern(pattern).Match(NewBindings(), value, nil); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[E=(quote yes) @[{1 2}]][P=(quote t) @[{1 2}]][E=(quote no) @[{0 2}]][P=(quote ()) @[{0 2}]]}", bindings.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
