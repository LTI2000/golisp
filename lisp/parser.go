package lisp

import (
	"errors"
)

type Parser struct {
	tokenizer  *Tokenizer
	tokenStack []*Token
}

func (p *Parser) nextToken() (*Token, error) {
	if token := p.popToken(); token != nil {
		return token, nil
	} else {
		return p.tokenizer.NextToken()
	}
}

func (p *Parser) popToken() *Token {
	pushbackCount := len(p.tokenStack)
	if pushbackCount > 0 {
		token := p.tokenStack[pushbackCount-1]
		p.tokenStack[pushbackCount-1] = nil
		p.tokenStack = p.tokenStack[:pushbackCount-1]
		return token
	} else {
		return nil
	}
}

func (p *Parser) pushToken(token *Token) {
	p.tokenStack = append(p.tokenStack, token)
}

func NewParser(t *Tokenizer) *Parser {
	return &Parser{t, nil}
}

func (p *Parser) ParseExpression() (Value, error) {
	return p.parseExpression()
}

func (p *Parser) parseExpression() (Value, error) {
	if token, err := p.nextToken(); err != nil {
		return nil, err
	} else if token.Type == Identifier {
		return Symbol(token.Value), nil
	} else if token.Type == Apostrophe {
		if exp, err := p.parseExpression(); err != nil {
			return nil, err
		} else {
			return List(Quote, exp), nil
		}
	} else if token.Type == LeftParen {
		return p.parseList()
	} else {
		return nil, errors.New("illegal expression")
	}
}

func (p *Parser) parseList() (Value, error) {
	if token, err := p.nextToken(); err != nil {
		return nil, err
	} else if token.Type == RightParen {
		return Nil, nil
	} else {
		p.pushToken(token)
	}

	car, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if token, err := p.nextToken(); err != nil {
		return nil, err
	} else if token.Type == Dot {
		cdr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		return Cons(car, cdr), nil

	} else {
		p.pushToken(token)
		cdr, err := p.parseList()
		if err != nil {
			return nil, err
		}
		return Cons(car, cdr), nil
	}
}
