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
	Identifier
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
	} else if unicode.IsLetter(char) {
		t.buffer = append([]rune(nil), char)
		if err := readIdentifier(t); err != nil {
			return nil, err
		}
		return &Token{Identifier, string(t.buffer)}, nil
	} else {
		return nil, errors.New("illegal token")
	}
}

func skipSpace(t *Tokenizer) error {
	for {
		if char, _, err := t.reader.ReadRune(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			} else {
				return err
			}
		} else if !unicode.IsSpace(char) {
			return t.reader.UnreadRune()
		}
	}
}

func readIdentifier(t *Tokenizer) error {
	for {
		if char, _, err := t.reader.ReadRune(); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			} else {
				return err
			}
		} else if unicode.IsLetter(char) || unicode.IsNumber(char) {
			t.buffer = append(t.buffer, char)
		} else {
			return t.reader.UnreadRune()
		}
	}
}
