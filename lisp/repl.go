package lisp

import (
	"fmt"
)

func Repl(reader *Reader, env Environment) Environment {
	for {
		if expression, err := reader.ReadValue(); err != nil {
			fmt.Printf("read failed: %v\n", err.Error())
			return env
		} else if expression == nil {
			return env
		} else if result, bindings, err := evalTopLevel(expression, env); err != nil {
			fmt.Printf("invalid top level expression: %v\n", err.Error())
			return env
		} else {
			env = bindings
			if result != nil {
				fmt.Printf("; %v\n", result)
			}
		}
	}
}

// handles defun expressions, as well as general expression evaluation
func evalTopLevel(exp Expression, env Environment) (Expression, Environment, error) {
	if name, args, body, ok := Match3("(defun NAME ARGS BODY)", exp, "NAME", "ARGS", "BODY"); ok {
		return nil, Extend(name.String(), List(LABEL, name, List(LAMBDA, args, body)), env), nil
	} else {
		if result, err := Eval(exp, env); err != nil {
			return nil, nil, err
		} else {
			return result, env, nil
		}
	}
}
