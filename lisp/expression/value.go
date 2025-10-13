package expression

import (
	"fmt"
	"strings"
)

// Value
type Value interface {
	prim_atom() bool
	prim_car() (Value, error)
	prim_cdr() (Value, error)
	prim_eq(Value) bool
	String() string
}

func Atom(x Value) bool {
	return x.prim_atom()
}

func Car(x Value) (Value, error) {
	return x.prim_car()
}

func Cdr(x Value) (Value, error) {
	return x.prim_cdr()
}
func Eq(x, y Value) bool {
	return x.prim_eq(y)
}

var T Value = Symbol("t")
var Nil Value = Symbol("nil")
var Quote Value = Symbol("quote")

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

func (*symbol) prim_atom() bool {
	return true
}

func (s *symbol) prim_car() (Value, error) {
	return nil, fmt.Errorf("car: not a cons: %v", s)
}

func (s *symbol) prim_cdr() (Value, error) {
	return nil, fmt.Errorf("cdr: not a cons: %v", s)
}

func (s *symbol) prim_eq(other Value) bool {
	switch v := other.(type) {
	case *symbol:
		return s.name == v.name
	default:
		return false
	}
}

func (s *symbol) String() string {
	if s == Nil {
		return "()"
	} else {
		return s.name
	}
}

// Pair
type cons struct {
	car Value
	cdr Value
}

func Cons(car Value, cdr Value) Value {
	return &cons{car, cdr}
}

func (*cons) prim_atom() bool {
	return false
}

func (c *cons) prim_car() (Value, error) {
	return c.car, nil
}

func (c *cons) prim_cdr() (Value, error) {
	return c.cdr, nil
}
func (c *cons) prim_eq(other Value) bool {
	return false
}

func (c *cons) String() string {
	var sb strings.Builder
	sb.WriteString("(")

	sb.WriteString(c.car.String())
	rest := c.cdr
loop:
	for {
		switch v := rest.(type) {
		case *cons:
			sb.WriteString(" ")
			sb.WriteString(v.car.String())
			rest = v.cdr
		case *symbol:
			if v != Nil {
				sb.WriteString(" . ")
				sb.WriteString(v.String())
			}
			break loop
		}
	}

	sb.WriteString(")")
	return sb.String()
}

// Utils

// create a (possibly empty) List from a Value slice.
func List(slice ...Value) Value {
	list := Nil
	for i := len(slice) - 1; i >= 0; i-- {
		list = Cons(slice[i], list)
	}
	return list
}

// // create a Value slice from a Value, which must be a list. panics if not.
// func Slice(list Value) (slice []Value) {
// 	for list != Nil {
// 		slice = append(slice, list.GetCar())
// 		list = list.GetCdr()
// 	}
// 	return
// }

// // append two lists. Panics if l1 or l2 is not a list Value
// func Append(l1, l2 Value) Value {
// 	if l1 != Nil {
// 		return Cons(l1.GetCar(), Append(l1.GetCdr(), l2))
// 	} else {
// 		return l2
// 	}
// }

// func Foldr(f func(Value, Value) Value, z, l Value) Value {
// 	if l != Nil {
// 		return f(l.GetCar(), Foldr(f, z, l.GetCdr()))
// 	} else {
// 		return z
// 	}
// }

// func Concat(l Value) Value {
// 	return Foldr(Append, Nil, l)
// }

// // concatMap               :: (a -> [b]) -> [a] -> [b]
// // concatMap f             =  foldr ((++) . f) []
