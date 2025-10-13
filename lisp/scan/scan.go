package scan

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

func (t *Token) String() string {
	return t.Value
}

type Scanner struct {
	reader bufio.Reader
	buffer []rune
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{*bufio.NewReader(reader), make([]rune, 0, 16)}
}

func (s *Scanner) NextToken() (*Token, error) {
	if err := skipSpace(s); err != nil {
		return nil, err
	}

	if char, _, err := s.reader.ReadRune(); err != nil {
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
		s.buffer = append([]rune(nil), char)
		if err := readIdentifier(s); err != nil {
			return nil, err
		}
		name := string(s.buffer)
		if name == "." {
			return &Token{Dot, name}, nil
		} else {
			return &Token{Identifier, name}, nil
		}
	} else {
		return nil, errors.New("illegal token")
	}
}

func skipSpace(s *Scanner) error {
	for {
		if char, _, err := s.reader.ReadRune(); err != nil {
			return maskEof(err)
		} else if !unicode.IsSpace(char) {
			return s.reader.UnreadRune()
		}
	}
}

func readIdentifier(s *Scanner) error {
	for {
		if char, _, err := s.reader.ReadRune(); err != nil {
			return maskEof(err)
		} else if isIdentifierChar(char) || unicode.IsNumber(char) {
			s.buffer = append(s.buffer, char)
		} else {
			return s.reader.UnreadRune()
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
