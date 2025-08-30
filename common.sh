#!/bin/bash

echo "creating base files.. ${1}"

file="$1/${1}.go"
file_test="$1/${1}_test.go"

mkdir "$1"
touch $file
touch ${file_test}
echo -e "package $1\n\n/* Key ideas:\n*\n*\n */" >$file
echo -e "package $1\n\n/* Key ideas:\n*\n*\n */" >$file_test
