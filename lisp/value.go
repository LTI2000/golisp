package lisp

// Value
type Value interface {
	Atom() bool
	Car() Value
	Cdr() Value
	Eq(other Value) bool
}

// Symbol
type symbol struct {
	name string
}

var symbols map[string]*symbol = make(map[string]*symbol)

func Symbol(name string) Value {
	value, ok := symbols[name]
	if !ok {
		value = &symbol{name}
		symbols[name] = value
	}
	return value
}

func (*symbol) Atom() bool {
	return true
}

func (*symbol) Car() Value {
	panic("car: got symbol")
}

func (*symbol) Cdr() Value {
	panic("cdr: got symbol")
}

func (sym *symbol) Eq(other Value) bool {
	switch v := other.(type) {
	case *symbol:
		return sym.name == v.name
	default:
		return false
	}
}

// Cons
type cons struct {
	car Value
	cdr Value
}

func Pair(car Value, cdr Value) Value {
	return &cons{car, cdr}
}

func (*cons) Atom() bool {
	return false
}

func (pair *cons) Car() Value {
	return pair.car
}

func (pair *cons) Cdr() Value {
	return pair.cdr
}
func (pair *cons) Eq(other Value) bool {
	return false
}
