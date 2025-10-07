package lisp

import (
	"errors"
)

type Reader struct {
	tokenizer  *Tokenizer
	tokenStack []*Token
}

func (r *Reader) nextToken() (*Token, error) {
	if token := r.popToken(); token != nil {
		return token, nil
	} else {
		return r.tokenizer.NextToken()
	}
}

func (r *Reader) popToken() *Token {
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

func (r *Reader) pushToken(token *Token) {
	r.tokenStack = append(r.tokenStack, token)
}

func (r *Reader) peekToken(tokenType TokenType) (bool, error) {
	if token, err := r.nextToken(); err != nil {
		return false, err
	} else if token.Type == tokenType {
		return true, nil
	} else {
		r.pushToken(token)
		return false, nil
	}
}

func NewReader(t *Tokenizer) *Reader {
	return &Reader{t, nil}
}

func (r *Reader) ReadExpression() (Value, error) {
	return r.readExpression()
}

func (r *Reader) readExpression() (Value, error) {
	if token, err := r.nextToken(); err != nil {
		return nil, err
	} else if token.Type == Identifier {
		return Symbol(token.Value), nil
	} else if token.Type == Apostrophe {
		if exp, err := r.readExpression(); err != nil {
			return nil, err
		} else {
			return List(Quote, exp), nil
		}
	} else if token.Type == LeftParen {
		return r.readList()
	} else {
		return nil, errors.New("illegal expression")
	}
}

func (r *Reader) readList() (Value, error) {
	if isRightParen, err := r.peekToken(RightParen); err != nil {
		return nil, err
	} else if isRightParen {
		return Nil, nil
	} else if head, err := r.readExpression(); err != nil {
		return nil, err
	} else if isDot, err := r.peekToken(Dot); err != nil {
		return nil, err
	} else if tail, err := r.readTail(isDot); err != nil {
		return nil, err
	} else {
		return Pair(head, tail), nil
	}
}

func (r *Reader) readTail(isDot bool) (tail Value, err error) {
	if isDot {
		tail, err = r.readExpression()
	} else {
		tail, err = r.readList()
	}
	return
}
