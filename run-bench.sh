#!/bin/sh
# ==============================================================================
#  Benchmark Script to measure the performance of the Hash function
# ==============================================================================
#  This script takes aroung 15-20 minutes to run under:
#    cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
# ==============================================================================

path_bench="./bench_results"
name_file="bench_results"
path_file="${path_bench}/${name_file}"

set -e

# ------------------------------------------------------------------------------
#  Benchmarking
# ------------------------------------------------------------------------------
#  The below runs the benchmark 5 times for 5 seconds, each time with a different
#  order. The results will be sorted in the next step.
# ------------------------------------------------------------------------------
go clean -testcache # clean to expire all cached test results

echo "- Benchmarking  ..."
echo '  Settings: -failfast -benchmem -benchtime=5s -count=1- shuffle=on (x5 times)'

go test -failfast -benchmem -benchtime=5s -count=1 -bench . | tee "${path_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${path_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${path_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${path_file}.txt"
go test -failfast -benchmem -benchtime=5s -count=1 -shuffle=on -bench . | grep -e "^Benchmark" | tee -a "${path_file}.txt"

# ------------------------------------------------------------------------------
#  Sort the results by iter and time.
# ------------------------------------------------------------------------------
#  The following scripts require the installation of the below go packages:
#    go install github.com/KEINOS/go-prettybench@latest
#    go install golang.org/x/perf/cmd/benchstat@latest
# ------------------------------------------------------------------------------
type go-prettybench >/dev/null || {
    echo "go-prettybench is not installed. The outputs are not sorted."
    echo "Bench result file: ${path_file}.txt"
    exit
}

type benchstat >/dev/null || {
    echo "benchstat is not installed. Benchmarks statistics were not computed and compared."
    echo "Bench result file: ${path_file}.txt"
    exit
}

echo "- Sorting benchmark results ..."
go-prettybench -sort iter >"${path_file}_sorted.txt" <"${path_file}.txt"
benchstat "${path_file}_sorted.txt" >"${path_file}_stats.txt"

echo "Sorted by iter: ${path_file}_sorted.txt"
echo "Statistics    : ${path_file}_stats.txt"
