package lisp

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokenizer *Tokenizer
	lookahead *Token
}

func NewParser(t *Tokenizer) *Parser {
	return &Parser{t, nil}
}

func (p *Parser) ParseExpression() (Value, error) {
	if err := p.lexan(); err != nil {
		return nil, err
	} else {
		return p.parseExpression()
	}
}

func (p *Parser) lexan() error {
	if lookahead, err := p.tokenizer.NextToken(); err != nil {
		return err
	} else {
		p.lookahead = lookahead
		fmt.Printf("lookahead %v\n", p.lookahead)
		return nil
	}
}

func (p *Parser) parseExpression() (Value, error) {
	if p.lookahead.Type == Identifier {
		if sym, err := Symbol(p.lookahead.Value), p.lexan(); err != nil {
			return nil, err
		} else {
			return sym, nil
		}
	} else if p.lookahead.Type == LeftParen {
		if err := p.lexan(); err != nil {
			return nil, err
		} else {
			return p.parseList()
		}
	} else {
		return nil, errors.New("illegal expression")
	}
}
func (p *Parser) parseList() (Value, error) {
	if p.lookahead.Type == RightParen {
		if err := p.lexan(); err != nil {
			return nil, err
		} else {
			return Symbol("nil"), nil
		}
	} else if car, err := p.parseExpression(); err != nil {
		return nil, err
	} else if cdr, err := p.parseList(); err != nil {
		return nil, err
	} else {
		return Cons(car, cdr), nil
	}
}
