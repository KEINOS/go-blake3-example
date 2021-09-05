#!/bin/sh

name_file="bench_results"

set -e

go clean
go test -failfast -benchmem -benchtime=5s -count=1 -bench . | tee "${name_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${name_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${name_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${name_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${name_file}.txt"

#clear
#echo "Sorted Benchmark Results"
#cat "$name_file.txt" | go-prettybench -sort iter > "${name_file}_sorted.txt"
#benchstat "${name_file}_sorted.txt" > "${name_file}_stats.txt"
