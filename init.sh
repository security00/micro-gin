#!/bin/bash

dirs=`find . -type d  | grep -v git`

for dir in ${dirs}
do
touch ${dir}/delete.txt
done

