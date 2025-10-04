package lisp

import (
	"bufio"
	"io"
	"unicode"
)

type TokenType int

const (
	LeftParen TokenType = iota
	RightParen
)

type Token struct {
	Type  TokenType
	Value string
}
type Tokenizer struct {
	reader bufio.Reader
	buffer []rune
}

func NewTokenizer(reader io.Reader) *Tokenizer {
	return &Tokenizer{*bufio.NewReader(reader), make([]rune, 0, 16)}
}

func (t *Tokenizer) NextToken() (*Token, error) {
	err := skipSpace(t)
	if err != nil {
		return nil, err
	}
	char, _, err := t.reader.ReadRune()
	if err != nil {
		return nil, err
	}
	switch char {
	case '(':
		return &Token{LeftParen, "("}, nil
	case ')':
		return &Token{RightParen, ")"}, nil
	}
	return nil, nil
}

func skipSpace(t *Tokenizer) error {
	for {
		char, _, err := t.reader.ReadRune()
		if err != nil {
			return err
		}
		if !unicode.IsSpace(char) {
			err := t.reader.UnreadRune()
			if err != nil {
				return err
			}
			return nil
		}
	}
}
