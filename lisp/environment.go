package lisp

import (
	"fmt"
	"strings"
)

type Environment interface {
	Lookup(name Expression) (Expression, error)
	String() string
}

func NewEnvironment() Environment {
	return &empty_env{}
}

func Extend(name Expression, value Expression, env Environment) Environment {
	return &extended_env{name, value, env}
}

func ExtendList(names []Expression, values []Expression, env Environment) Environment {
	nameCount := len(names)
	valueCount := len(values)
	if nameCount != valueCount {
		panic("name/value count mismatch")
	} else {
		for i := nameCount - 1; i >= 0; i-- {
			env = Extend(names[i], values[i], env)
		}
		return env
	}
}

func Merge(e1, e2 Environment) Environment {
	switch e := e1.(type) {
	case *extended_env:
		return Extend(e.name, e.value, Merge(e.next, e2))
	default:
		return e2
	}
}

//

type empty_env struct{}

func (e *empty_env) Lookup(name Expression) (Expression, error) {
	return nil, fmt.Errorf("%v is unbound", name)
}

func (e *empty_env) String() string {
	return "[]"
}

//

type extended_env struct {
	name  Expression
	value Expression
	next  Environment
}

func (e *extended_env) Lookup(name Expression) (Expression, error) {
	if name == e.name {
		return e.value, nil
	} else {
		return e.next.Lookup(name)
	}

}

func (e *extended_env) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	sb.WriteString(e.name.String())
	sb.WriteString(" := ")
	sb.WriteString(e.value.String())
	rest := e.next
loop:
	for {
		switch e := rest.(type) {
		case *extended_env:
			sb.WriteString(", ")
			sb.WriteString(e.name.String())
			sb.WriteString(" := ")
			sb.WriteString(e.value.String())
			rest = e.next
		default:
			break loop
		}
	}

	sb.WriteString("]")
	return sb.String()
}
