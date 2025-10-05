package lisp

import (
	"strings"
	"testing"
)

func TestParseExpressionSymbol(t *testing.T) {
	reader := strings.NewReader("quote")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if actual, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected := Symbol("quote"); actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseExpressionEmptyList(t *testing.T) {
	reader := strings.NewReader("()")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if actual, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected := Symbol("nil"); actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseExpressionSigletonList(t *testing.T) {
	reader := strings.NewReader("(x)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if actual, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if Atom(actual) != Symbol("nil") {
		t.Errorf("expected list")
	} else if car := Car(actual); car != Symbol("x") {
		t.Errorf("expected symbol x")
	} else if cdr := Cdr(actual); cdr != Symbol("nil") {
		t.Errorf("expected empty list")
	}
}
func TestParseExpressionProperList(t *testing.T) {
	reader := strings.NewReader("(x y)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if actual, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if Atom(actual) != Symbol("nil") {
		t.Errorf("expected list")
	} else if car := Car(actual); car != Symbol("x") {
		t.Errorf("expected symbol x")
	} else if cdr := Cdr(actual); cdr == Symbol("nil") {
		t.Errorf("expected list")
	} else if cadr := Car(Cdr(actual)); cadr != Symbol("y") {
		t.Errorf("expected symbol y")
	} else if cddr := Cdr(Cdr(actual)); cddr != Symbol("nil") {
		t.Errorf("expected empty list")
	}
}
