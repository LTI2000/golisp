package lisp

import "testing"

func TestMatchSymbol(t *testing.T) {
	pattern := Must(Read, "()")
	value := Must(Read, "nil")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchPair(t *testing.T) {
	pattern := Must(Read, "(a b)")
	value := Must(Read, "(a . (b . ()))")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariable(t *testing.T) {
	pattern := Must(Read, "X")
	value := Must(Read, "t")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[X=t]}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair1(t *testing.T) {
	pattern := Must(Read, "(X)")
	value := Must(Read, "t")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); matches {
		t.Errorf("expected non-match: %v %v", pattern, value)
	} else if expected, actual := "{}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchVariablePair2(t *testing.T) {
	pattern := Must(Read, "(X Y)")
	value := Must(Read, "(y x)")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[Y=x][X=y]}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMatchCond(t *testing.T) {
	pattern := Must(Read, "(cond (P E) ...)")
	value := Must(Read, "(cond ('() 'no) ('t 'yes))")

	if bindings, matches := NewPattern(pattern).Match(nil, value, false); !matches {
		t.Errorf("expected match: %v %v", pattern, value)
	} else if expected, actual := "{[P=((quote ()) (quote t))][E=((quote no) (quote yes))]}", String(bindings); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
