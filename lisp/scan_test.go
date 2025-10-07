package lisp

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	reader := strings.NewReader(" (\t foo123 bar baz.)\n")
	tokenizer := NewScanner(reader)

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := LeftParen, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "foo123", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "bar", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Identifier, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	} else if expected, actual := "baz.", token.Value; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := RightParen, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestDot(t *testing.T) {
	reader := strings.NewReader(".")
	tokenizer := NewScanner(reader)

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Dot, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestEmptyInput(t *testing.T) {
	reader := strings.NewReader("")
	tokenizer := NewScanner(reader)

	if token, err := tokenizer.NextToken(); err != nil {
		t.Fatalf("err %v", err)
	} else if expected, actual := Eof, token.Type; expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
