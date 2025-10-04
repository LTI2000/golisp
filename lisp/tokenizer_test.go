package lisp

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	reader := strings.NewReader(" (\t)\n")
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
	if token.Type != RightParen {
		t.Errorf("expected )")
	}
}
