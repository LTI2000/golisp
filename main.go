package main

import (
	"fmt"
	"io"
	"os"

	"github.com/LTI2000/golisp/lisp"
)

func main() {
	fmt.Printf("; welcome to golisp\n")
	repl(os.Stdin)
}

func repl(r io.Reader) {
	scanner := lisp.NewScanner(r)
	reader := lisp.NewReader(scanner)
	for {
		if value, err := reader.ReadValue(); err != nil {
			fmt.Printf("read failed: %v\n", err.Error())
			return
		} else if expression, err := lisp.ParseExpression(value); err != nil {
			fmt.Printf("parse failed: %v\n", err.Error())
			return
		} else if v, err := lisp.Eval(expression); err != nil {
			fmt.Printf("eval failed: %v\n", err.Error())
			return
		} else {
			fmt.Printf("; %v\n", v)
		}
	}
}
