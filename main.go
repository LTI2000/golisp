package main

import (
	"fmt"
	"io"
	"os"

	"github.com/LTI2000/golisp/lisp"
	"github.com/LTI2000/golisp/lisp/scan"
)

func main() {
	fmt.Printf(";; welcome to golisp\n")
	repl(os.Stdin)
	fmt.Printf(";; repl finished\n")
}

func repl(r io.Reader) {
	scanner := scan.NewScanner(r)
	reader := lisp.NewReader(scanner)

	lisp.Repl(reader)
}
