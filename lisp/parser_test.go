package lisp

import (
	"strings"
	"testing"
)

func TestParseSymbol(t *testing.T) {
	reader := strings.NewReader("quote")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("quote"), expression; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseQuotation(t *testing.T) {
	reader := strings.NewReader("'a")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Quote, expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("a"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseEmptyList(t *testing.T) {
	reader := strings.NewReader("()")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Nil, expression; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseSingletonList(t *testing.T) {
	reader := strings.NewReader("(x)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseProperList(t *testing.T) {
	reader := strings.NewReader("(x y)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseNestedList(t *testing.T) {
	reader := strings.NewReader("(x (y) z)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCar().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("z"), expression.GetCdr().GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Nil, expression.GetCdr().GetCdr().GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseDottedList(t *testing.T) {
	reader := strings.NewReader("(x . y)")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestParseDottedProperList(t *testing.T) {
	reader := strings.NewReader("(x . (y . ())")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Symbol("x"), expression.GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := Symbol("y"), expression.GetCdr().GetCar(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
func TestParseLongDottedProperList(t *testing.T) {
	reader := strings.NewReader("(x . (y . (z . ()))")
	tokenizer := NewTokenizer(reader)
	parser := NewParser(tokenizer)

	if expression, err := parser.ParseExpression(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := "(x y z)", expression.String(); expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
