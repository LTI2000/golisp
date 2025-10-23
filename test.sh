go clean
go build
go test ./...
cat test.lisp eval.lisp eval_test.lisp | go run main.go > actual.txt
diff expected.txt actual.txt
