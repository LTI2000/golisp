package lisp

type Pattern interface {
	Match(Value) bool
}

type symbolPattern struct {
	symbol Value
}

func (p *symbolPattern) Match(value Value) bool {
	return p.symbol.IsEq(value)
}

type variablePattern struct {
	name Value
}

func (p *variablePattern) Match(Value) bool {
	return false
}

type pairPattern struct {
	head Pattern
	tail Pattern
}

func (p *pairPattern) Match(Value) bool {
	return false
}

func NewPattern(pattern Value) Pattern {
	return &symbolPattern{pattern}
}
