package lisp

import (
	"fmt"
)

func Repl(reader *Reader) {
	env := NIL
	for {
		if expression, err := reader.ReadValue(); err != nil {
			fmt.Printf("read failed: %v\n", err.Error())
			return
		} else if expression == nil {
			return
		} else if result, bindings, err := evalTopLevel(expression, env); err != nil {
			fmt.Printf("invalid top level expression: %v\n", err.Error())
			return
		} else {
			env = bindings
			if result != nil {
				fmt.Printf("; %v\n", result)
			}
		}
	}
}

// handles defun expressions, as well as general expression evaluation
func evalTopLevel(e Expression, a Expression) (Expression, Expression, error) {
	if name, args, body, ok := Match3("(defun NAME ARGS BODY)", e, "NAME", "ARGS", "BODY"); ok {
		return nil, Must2(Append, List(List(name, List(LABEL, name, List(LAMBDA, args, body)))), a), nil
	} else {
		if result, err := Eval(e, a); err != nil {
			return nil, nil, err
		} else {
			return result, a, nil
		}
	}
}
