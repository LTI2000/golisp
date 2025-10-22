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
		} else if result, bindings, err := TopLevel(expression, env); err != nil {
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
