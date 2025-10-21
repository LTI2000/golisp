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
		} else if binding, err := Defun(expression); err != nil {
			//fmt.Printf("defun failed: %v\n", err.Error())
			if result, err := Eval(expression, env); err != nil {
				fmt.Printf("eval failed: %v\n", err.Error())
				return
			} else {
				fmt.Printf("; %v\n", result)
			}
		} else {
			env = Cons(binding, env)
			fmt.Printf("binding: %v\n", binding)
		}
	}
}
