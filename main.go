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
	repl0(os.Stdin)
}

func repl0(r io.Reader) {
	scanner := scan.NewScanner(r)
	reader := lisp.NewReader(scanner)

	lisp.Repl(reader)
}
