# Example Usage of [BLAKE3 for Go](https://github.com/lukechampine/blake3)

Example usge and benchmark results of [BLAKE3](https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE3) implementation for Go by [@lukechampine](https://github.com/lukechampine/blake3).

```go
// import "lukechampine.com/blake3"
input := "This is a string"

valByte := blake3.Sum256([]byte(input))

// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2"
actual := fmt.Sprintf("%x", valByte)

fmt.Println("expect:", expect)
fmt.Println("actual:", actual)
// Output:
// expect: 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
// actual: 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
```

- About BLAKE3
    - [BLAKE3](https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE3) | BLAKE(hash function) @ Wikipedia
    - [Official repo for Rust and C](https://github.com/BLAKE3-team/BLAKE3) @ GitHub
- About BLAKE3 for Go
    - `go get lukechampine.com/blake3`
    - [Repo](https://github.com/lukechampine/blake3) | lukechampine @ GitHub
    - [Document](https://pkg.go.dev/lukechampine.com/blake3) @ pkg.go.dev

## Bench Results

OMG ... even `blake3.Sum512()`(64Bytes long hash) is faster than `fnv.New32()`(4Bytes long hash, [FNV1](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function)).

... Am I [doing something](./benchmark_test.go) wrong?

```shellsession
$ go test -failfast -benchmem -shuffle=on -benchtime=10s -count=2 -bench . | go-prettybench -sort iter
-test.shuffle 1630647622685374698
goos: linux
goarch: amd64
pkg: github.com/KEINOS/go-blake3-example
cpu: Intel(R) Xeon(R) Platinum 8168 CPU @ 2.70GHz
PASS
benchmark                                  iter         time/iter   bytes alloc        allocs
---------                                  ----         ---------   -----------        ------
BenchmarkCRC_32b-4                       244780       47.65 μs/op        8 B/op   1 allocs/op
BenchmarkCRC_32b-4                       243198       47.96 μs/op        8 B/op   1 allocs/op
BenchmarkMMH3_128b-4                      59395      202.11 μs/op       16 B/op   1 allocs/op
BenchmarkMMH3_128b-4                      58870      200.28 μs/op       16 B/op   1 allocs/op
BenchmarkBlake3_verLuke_512b-4            45450      267.98 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b-4            45418      263.65 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b-4            45156      265.94 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-4     45074      267.69 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-4     45058      264.52 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-4     45038      267.10 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-4     44686      265.45 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-4            44203      268.50 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b-4           40299      303.85 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b-4           39654      304.50 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-4    39444      305.27 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-4    39351      305.11 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-4    39244      304.92 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-4    39152      305.36 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-4           38756      314.02 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-4           38486      309.93 μs/op        0 B/op   0 allocs/op
BenchmarkAdler_32b-4                      31203      388.88 μs/op        8 B/op   1 allocs/op
BenchmarkAdler_32b-4                      31192      387.46 μs/op        8 B/op   1 allocs/op
BenchmarkBlake2b_512b-4                   10000     1012.27 μs/op       64 B/op   1 allocs/op
BenchmarkSHA1_160b-4                      10000     1018.50 μs/op       24 B/op   1 allocs/op
BenchmarkBlake2b_256b-4                   10000     1033.08 μs/op       32 B/op   1 allocs/op
BenchmarkBlake2b_256b-4                   10000     1024.48 μs/op       32 B/op   1 allocs/op
BenchmarkBlake2b_512b-4                   10000     1014.98 μs/op       64 B/op   1 allocs/op
BenchmarkSHA1_160b-4                      10000     1006.89 μs/op       24 B/op   1 allocs/op
BenchmarkMD5_128b-4                        8367     1424.84 μs/op       16 B/op   1 allocs/op
BenchmarkMD5_128b-4                        8065     1425.79 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_512b-4                        6822     1723.21 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_512b-4                        6532     1724.44 μs/op       64 B/op   1 allocs/op
BenchmarkFnv1a_128b-4                      5086     2378.60 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1a_128b-4                      5007     2381.48 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1_128b-4                       4794     2489.41 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_256b-4                        4762     2582.09 μs/op       32 B/op   1 allocs/op
BenchmarkFnv1_128b-4                       4761     2487.18 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_256b-4                        4658     2551.49 μs/op       32 B/op   1 allocs/op
BenchmarkSHA3_256b-4                       3835     3088.63 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_256b-4                       3625     3077.95 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_512b-4                       2097     5814.51 μs/op      576 B/op   3 allocs/op
BenchmarkSHA3_512b-4                       1995     5791.61 μs/op      576 B/op   3 allocs/op
BenchmarkWhirlpool_512b-4                   478    25161.96 μs/op       64 B/op   1 allocs/op
BenchmarkWhirlpool_512b-4                   475    25015.15 μs/op       64 B/op   1 allocs/op
ok  	github.com/KEINOS/go-blake3-example	582.746s
```

- This repo is GitHub Codespaces compatible. Try your own online on GitHub.
