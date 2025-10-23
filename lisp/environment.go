package lisp

import "fmt"

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
	panic("unimplemented")
}
