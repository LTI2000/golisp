package main

import (
	"fmt"
	"io"
	"os"

	"github.com/LTI2000/golisp/lisp"
	"github.com/LTI2000/golisp/lisp/scan"
)

func main() {
	fmt.Printf("; welcome to golisp\n")
	repl(os.Stdin)
}

func repl(r io.Reader) {
	scanner := scan.NewScanner(r)
	reader := lisp.NewReader(scanner)
	for {
		if expression, err := reader.ReadValue(); err != nil {
			fmt.Printf("read failed: %v\n", err.Error())
			return
		} else {
			fmt.Printf("; %v\n", expression)
		}
	}
}
