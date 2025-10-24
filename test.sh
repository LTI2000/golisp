#!/bin/bash
echo ""
echo "################"
echo "CLEAN BUILD TEST"
echo "################"
go clean
go build
go test ./...
echo "################"
echo "DONE"
echo "################"
echo ""

echo ""
echo "################"
echo "END TO END TEST"
echo "################"
cat lisp-src/test/test.lisp lisp-src/test/eval_test.lisp | go run main.go > actual.txt

diff -q expected.txt actual.txt
status=$?
if [ $status -gt 0 ]; then
  diff -y --color=always expected.txt actual.txt
  echo "################"
  echo -e "\033[0;31m*** FAIL ***\033[0m"
  echo "################"
else
  echo "################"
  echo -e "\033[0;32m**** OK ****\033[0m"
  echo "################"
fi
echo ""
