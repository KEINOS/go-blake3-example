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
$ go test -failfast -benchmem -shuffle=on -benchtime=10s -count=5 -bench . | go-prettybench -sort iter
-test.shuffle 1630567361896437483
goos: linux
goarch: amd64
pkg: github.com/KEINOS/go-blake3-example
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
PASS
benchmark                   iter        time/iter   bytes alloc        allocs
---------                   ----        ---------   -----------        ------
BenchmarkBlake3_64Byte-2   21663     540.50 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   21351     562.46 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   21298     544.14 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   21201     541.83 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   21181     569.72 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   21098     542.34 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   20890     560.94 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   20629     539.04 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   20097     560.38 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   19732     689.09 μs/op        0 B/op   0 allocs/op
BenchmarkFNV1A_4Byte-2      8366    1409.17 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       8178    1413.68 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      8164    1432.98 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       8056    1409.29 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       8022    1397.08 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       7976    1400.80 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       7804    1419.16 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      7752    1404.71 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      7546    1411.91 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      7362    1428.08 μs/op        8 B/op   1 allocs/op
BenchmarkSHA3_32Byte-2      2994    3817.32 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2932    3835.08 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2896    3869.84 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2788    3794.07 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2676    4052.84 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1630    7323.89 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1608    7243.38 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1606    7214.51 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1509    7304.93 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1500    7214.14 μs/op     1024 B/op   4 allocs/op
ok      github.com/KEINOS/go-blake3-example     408.685s
```

- This repo is GitHub Codespaces compatible. Try your own online on GitHub.
