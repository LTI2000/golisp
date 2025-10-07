package lisp

import (
	"bufio"
	"errors"
	"io"
	"unicode"
)

type TokenType int

const (
	LeftParen TokenType = iota
	RightParen
	Apostrophe
	Identifier
	Dot
	Eof
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
	if err := skipSpace(t); err != nil {
		return nil, err
	}

	if char, _, err := t.reader.ReadRune(); err != nil {
		if errors.Is(err, io.EOF) {
			return &Token{Eof, ""}, nil
		} else {
			return nil, err
		}
	} else if char == '(' {
		return &Token{LeftParen, string(char)}, nil
	} else if char == ')' {
		return &Token{RightParen, string(char)}, nil
	} else if char == '\'' {
		return &Token{Apostrophe, string(char)}, nil
	} else if isIdentifierChar(char) {
		t.buffer = append([]rune(nil), char)
		if err := readIdentifier(t); err != nil {
			return nil, err
		}
		name := string(t.buffer)
		if name == "." {
			return &Token{Dot, name}, nil
		} else {
			return &Token{Identifier, name}, nil
		}
	} else {
		return nil, errors.New("illegal token")
	}
}

func skipSpace(t *Tokenizer) error {
	for {
		if char, _, err := t.reader.ReadRune(); err != nil {
			return maskEof(err)
		} else if !unicode.IsSpace(char) {
			return t.reader.UnreadRune()
		}
	}
}

func readIdentifier(t *Tokenizer) error {
	for {
		if char, _, err := t.reader.ReadRune(); err != nil {
			return maskEof(err)
		} else if isIdentifierChar(char) || unicode.IsNumber(char) {
			t.buffer = append(t.buffer, char)
		} else {
			return t.reader.UnreadRune()
		}
	}
}

func isIdentifierChar(char rune) bool {
	return char == '.' || unicode.IsLetter(char)
}

func maskEof(err error) error {
	if errors.Is(err, io.EOF) {
		return nil
	} else {
		return err
	}
}
