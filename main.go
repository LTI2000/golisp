package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/LTI2000/golisp/lisp"
	"github.com/LTI2000/golisp/lisp/scan"
)

//go:embed lisp-src/main/*
var lisp_src embed.FS

func main() {
	env := lisp.NewEnvironment()

	lisp_fs, _ := fs.Sub(lisp_src, "lisp-src/main")
	fs.WalkDir(lisp_fs, ".",
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				if file, err := lisp_fs.Open(path); err != nil {

				} else {
					fmt.Printf("loading %v\n", path)
					env = repl(file, env)
				}
			}
			return nil
		})

	fmt.Printf(";; welcome to golisp\n")
	_ = repl(os.Stdin, env)
	fmt.Printf(";; repl finished\n")

}

func repl(r io.Reader, env lisp.Environment) lisp.Environment {
	scanner := scan.NewScanner(r)
	reader := lisp.NewReader(scanner)

	return lisp.Repl(reader, env)
}
