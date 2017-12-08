#!/bin/bash

grep '>' day7.txt | cut -d '>' -f 2 | tr ", " "\n" | grep -v "^$" | sort > day7-child.txt

cut -d ' ' -f 1 day7.txt | sort > day7-parents.txt

diff day7-parents.txt day7-child.txt