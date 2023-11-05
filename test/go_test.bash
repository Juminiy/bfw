#!/bin/bash

cd ../wheel
root_dir=$(pwd)
echo "cd "$(pwd)""

for dir in $(find . -type d)
do
    echo "cd "${root_dir}/${dir}" "

    cd "${root_dir}/${dir}"

    go test

    echo "${dir} test finished"
done
