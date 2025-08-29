#!/bin/bash

echo "create base files"
read -p "name: " name

file="$name/${name}.go"
file_test="$name/${name}_test.go"

mkdir "$name"
touch $file
touch ${file_test}
echo "package $name" >$file
echo "package $name" >$file_test
