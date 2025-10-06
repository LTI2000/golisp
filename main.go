package main

import (
	"github.com/LTI2000/golisp/lisp"
)

func main() {
	lisp.COND(lisp.List(lisp.Nil, lisp.T), lisp.List(lisp.T, lisp.Nil))
}
