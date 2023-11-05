#!/bin/bash

cd ../wheel
root_dir=$(pwd)
echo "cd "$(pwd)""

for dir in $(find . -type d)
do
    cd "${root_dir}/${dir}"

    find . -type f -perm +111 -exec file {} \; | grep ELF | cut -d: -f1 | xargs rm
done
