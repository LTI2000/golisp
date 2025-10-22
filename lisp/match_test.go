package lisp

import "testing"

func TestMatch0(t *testing.T) {
	if ok := Match0("x", Must(Read, "x")); !ok {
		t.Fatal("expected match")
	}

	if ok := Match0("x", Must(Read, "y")); ok {
		t.Fatal("expected no match")
	}

	if ok := Match0("x", Must(Read, "(x y)")); ok {
		t.Fatal("expected no match")
	}

	if ok := Match0("(x y)", Must(Read, "(x y)")); !ok {
		t.Fatal("expected match")
	}

	if ok := Match0("(x y)", Must(Read, "(x)")); ok {
		t.Fatal("expected no match")
	}

	if ok := Match0("(x y)", Must(Read, "(y x)")); ok {
		t.Fatal("expected no match")
	}
}

func TestMatch1(t *testing.T) {
	if value1, ok := Match1("X", Must(Read, "x"), "X"); !ok {
		t.Fatal("expected match")
	} else if expected, actual := "x", value1.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if value1, ok := Match1("X:atom", Must(Read, "x"), "X"); !ok {
		t.Fatal("expected match")
	} else if expected, actual := "x", value1.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if _, ok := Match1("X:atom", Must(Read, "(x y)"), "X"); ok {
		t.Fatal("expected no match")
	}
}

func TestMatch(t *testing.T) {
	for _, in := range [][3]string{
		{"(quote X)", "'x",
			"((X x))"},
		{"(atom X)", "(atom 'x)",
			"((X (quote x)))"},
		{"(eq X Y)", "(eq x y)",
			"((X x) (Y y))"},
		{"(car X)", "(car x)",
			"((X x))"},
		{"(cdr X)", "(cdr x)",
			"((X x))"},
		{"(cons X Y)", "(cons x y)",
			"((X x) (Y y))"},
		{"(cond . C)", "(cond (x 'nil) ('t 't))",
			"((C ((x (quote nil)) ((quote t) (quote t)))))"},
		{"(label F (lambda P E))", "(label id (lambda (x) x))",
			"((F id) (P (x)) (E x))"},
		{"((lambda P E) A)", "((lambda (x) x) 't)",
			"((P (x)) (E x) (A (quote t)))"},
		{"(defun F P E)", "(defun id (x) x)",
			"((F id) (P (x)) (E x))"},
		{"X", "x",
			"((X x))"},
		{"X:atom", "x",
			"((X x))"},
	} {
		pattern := Must(Read, in[0])
		expression := Must(Read, in[1])
		expected := Must(Read, in[2])
		if actual, ok :=
			match(pattern, expression); !ok || expected.String() != actual.String() {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}
}
