package lisp

import "errors"

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(t *Tokenizer) *Parser {
	return &Parser{t}
}

func (p *Parser) ParseExpression() (Value, error) {
	if token, err := p.tokenizer.NextToken(); err != nil {
		return nil, err
	} else if token.Type == Identifier {
		return Symbol(token.Value), nil
	} else if token.Type == LeftParen {
		return p.parseList()
	} else {
		return nil, errors.New("illegal expression")
	}
}

func (p *Parser) parseList() (Value, error) {
	if token, err := p.tokenizer.NextToken(); err != nil {
		return nil, err
	} else if token.Type == RightParen {
		return Symbol("nil"), nil
	} else if token.Type == Identifier {
		if rest, err := p.parseList(); err != nil {
			return nil, err
		} else {
			return Cons(Symbol(token.Value), rest), nil
		}
	} else {
		return nil, errors.New("illegal list expression")
	}
}
