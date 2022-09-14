#!/bin/bash

dirs=`find . -type d  | grep -v git`

for dir in ${dirs}
do
rm -rf ${dir}/delete.txt
done

