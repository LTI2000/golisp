package lisp

import "testing"

func TestRepl(t *testing.T) {
	Repl(StringReader(""), NewEnvironment())
}
