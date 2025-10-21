package lisp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/LTI2000/golisp/lisp/scan"
)

type Reader struct {
	scanner    *scan.Scanner
	tokenStack []*scan.Token
}

func NewReader(t *scan.Scanner) *Reader {
	return &Reader{t, nil}
}

func (r *Reader) nextToken() (*scan.Token, error) {
	if token := r.popToken(); token != nil {
		return token, nil
	} else {
		return r.scanner.NextToken()
	}
}

func (r *Reader) popToken() *scan.Token {
	pushbackCount := len(r.tokenStack)
	if pushbackCount > 0 {
		pushbackCount--
		token := r.tokenStack[pushbackCount]
		r.tokenStack[pushbackCount] = nil
		r.tokenStack = r.tokenStack[:pushbackCount]
		return token
	} else {
		return nil
	}
}

func (r *Reader) pushToken(token *scan.Token) {
	r.tokenStack = append(r.tokenStack, token)
}

func (r *Reader) peekToken(tokenType scan.TokenType) (bool, error) {
	if token, err := r.nextToken(); err != nil {
		return false, err
	} else if token.Type == tokenType {
		return true, nil
	} else {
		r.pushToken(token)
		return false, nil
	}
}

func (r *Reader) matchToken(token scan.TokenType) error {
	if match, err := r.peekToken(token); err != nil {
		return err
	} else if !match {
		return fmt.Errorf("missing token: '%v'", token)
	} else {
		return nil
	}
}

// returns nil, nil on eof
func (r *Reader) ReadValue() (Expression, error) {
	if token, err := r.nextToken(); err != nil {
		return nil, err
	} else if token.Type == scan.Eof {
		return nil, nil
	} else if token.Type == scan.Identifier {
		return Symbol(token.Value), nil
	} else if token.Type == scan.Apostrophe {
		if exp, err := r.ReadValue(); err != nil {
			return nil, err
		} else {
			return List(QUOTE, exp), nil
		}
	} else if token.Type == scan.LeftParen {
		return r.readList()
	} else {
		return nil, fmt.Errorf("illegal token: '%v'", token)
	}
}

func (r *Reader) readList() (Expression, error) {
	if isRightParen, err := r.peekToken(scan.RightParen); err != nil {
		return nil, err
	} else if isRightParen {
		return NIL, nil
	} else if head, err := r.ReadValue(); err != nil {
		return nil, err
	} else if head == nil {
		return nil, errors.New("unterminated list")
	} else if isDot, err := r.peekToken(scan.Dot); err != nil {
		return nil, err
	} else if tail, err := r.readTail(isDot); err != nil {
		return nil, err
	} else {
		return Cons(head, tail), nil
	}
}

func (r *Reader) readTail(isDot bool) (tail Expression, err error) {
	if isDot {
		tail, err = r.ReadValue()
		if err == nil {
			err = r.matchToken(scan.RightParen)
		}
		return
	} else {
		tail, err = r.readList()
	}
	return
}

// utility functions

func Read(source string) (Expression, error) {
	return StringReader(source).ReadValue()
}

func StringReader(source string) *Reader {
	return NewReader(scan.NewScanner(strings.NewReader(source)))
}
