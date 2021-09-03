#!/bin/sh

name_file="bench_results"

set -e

go clean
go test -failfast -benchmem -shuffle=on -benchtime=10s -count=2 -bench . | tee "${name_file}.txt"
clear
echo "Sorted Benchmark Results"
cat "$name_file.txt" | go-prettybench -sort iter > "${name_file}_sorted.txt"
