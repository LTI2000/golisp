#!/bin/bash
echo ""
echo "################"
echo "CLEAN BUILD TEST"
echo "################"
rm -f golisp
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
current_dir=..
pushd test
cat lisp-src/test.lisp lisp-src/eval_test.lisp | $current_dir/golisp > actual.txt
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
popd #test
echo ""
