#!/bin/sh

name_file="benchresults.txt"

set -e

go clean
go test -failfast -benchmem -shuffle=on -benchtime=10s -count=5 -bench . | tee "$name_file"
clear
echo "Sorted Benchmark Results"
cat "$name_file" | go-prettybench -sort iter
