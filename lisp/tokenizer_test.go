package lisp

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	reader := strings.NewReader(" (\t foo123 bar)\n")
	tokenizer := NewTokenizer(reader)

	token, err := tokenizer.NextToken()
	if err != nil {
		t.Fatalf("err %v", err)
	}
	if token.Type != LeftParen {
		t.Errorf("expected (")
	}

	token, err = tokenizer.NextToken()
	if err != nil {
		t.Fatalf("err %v", err)
	}
	if token.Type != Identifier {
		t.Errorf("expected Identifier")
	}
	if token.Value != "foo123" {
		t.Errorf("got %v, wanted %v", token.Value, "foo123")
	}

	token, err = tokenizer.NextToken()
	if err != nil {
		t.Fatalf("err %v", err)
	}
	if token.Type != Identifier {
		t.Errorf("expected Identifier")
	}
	if token.Value != "bar" {
		t.Errorf("got %v, wanted %v", token.Value, "bar")
	}

	token, err = tokenizer.NextToken()
	if err != nil {
		t.Fatalf("err %v", err)
	}
	if token.Type != RightParen {
		t.Errorf("expected )")
	}
}
