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
		pushbackCount--
		token := p.tokenStack[pushbackCount]
		p.tokenStack[pushbackCount] = nil
		p.tokenStack = p.tokenStack[:pushbackCount]
		return token
	} else {
		return nil
	}
}

func (p *Parser) pushToken(token *Token) {
	p.tokenStack = append(p.tokenStack, token)
}

func (p *Parser) peekToken(tokenType TokenType) (bool, error) {
	if token, err := p.nextToken(); err != nil {
		return false, err
	} else if token.Type == tokenType {
		return true, nil
	} else {
		p.pushToken(token)
		return false, nil
	}
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
	if isRightParen, err := p.peekToken(RightParen); err != nil {
		return nil, err
	} else if isRightParen {
		return Nil, nil
	} else if head, err := p.parseExpression(); err != nil {
		return nil, err
	} else if isDot, err := p.peekToken(Dot); err != nil {
		return nil, err
	} else if tail, err := p.parseTail(isDot); err != nil {
		return nil, err
	} else {
		return Pair(head, tail), nil
	}
}

func (p *Parser) parseTail(isDot bool) (tail Value, err error) {
	if isDot {
		tail, err = p.parseExpression()
	} else {
		tail, err = p.parseList()
	}
	return
}
