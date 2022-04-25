[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-blake3-example.svg)](https://pkg.go.dev/[github.com/KEINOS/go-pallet/pallet](https://pkg.go.dev/github.com/KEINOS/go-blake3-example))

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
  - [Official repo for Rust and C](https://github.com/BLAKE3-team/BLAKE3) | BLAKE3 Team @ GitHub
- About BLAKE3 for Go
  - We compared the below BLAKE3 implementations in Go. Both unofficial but works great.
  - Luke Champine's BLAKE3 for Go
    - `go get lukechampine.com/blake3`
    - [Repo](https://github.com/lukechampine/blake3) | lukechampine @ GitHub
    - [Document](https://pkg.go.dev/lukechampine.com/blake3) @ pkg.go.dev
  - Zeebo's BLAKE3 for Go
    - `go get github.com/zeebo/blake3`
    - [Repo](https://github.com/zeebo/blake3) | zeebo @ GitHub
    - [Document](https://pkg.go.dev/github.com/zeebo/blake3) @ pkg.go.dev

## Bench Results

O.M.G. ... ...

Even `blake3.Sum512()`(64 Bytes long BLAKE3 hash) is faster than `fnv.New32()`(4 Bytes long [FNV1](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function) hash).

Blake3-512bit (Luke ver.) was `676 μs/op ± 0%` and FNV1a-128bit was `3.53k μs/op ± 0%` ... ... Am I doing [something](./benchmark_test.go) wrong? If you have found something unfair, lease let us know in the [issues](https://github.com/KEINOS/go-blake3-example/issues). 😉

### Benchmarking and Statistics

```text
date: Mon, 25 Apr 2022 13:53:48 +0900
goos: darwin
goarch: amd64
pkg: github.com/KEINOS/go-blake3-example
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
tests: PASS

benchmark                                      iter         time/iter   bytes alloc        allocs
---------                                      ----         ---------   -----------        ------
BenchmarkCRC_32b-4                            87654       71.47 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit-4          66745       87.38 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-4   65907      108.02 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-4         57350      100.86 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-4    57160      105.10 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-4   27285      204.04 μs/op        0 B/op   0 allocs/op
BenchmarkMMH3_128b-4                          23210      253.23 μs/op       16 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-4         21585      292.43 μs/op        8 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b-4               13783      414.41 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-4        12142      488.15 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-4               10892      588.31 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-4         10843      597.17 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-4         10177      675.89 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-4        10094      562.13 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-4                 8479      727.34 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b-4                 8462      637.10 μs/op      256 B/op   1 allocs/op
BenchmarkAdler_32b-4                           8172      711.20 μs/op        8 B/op   1 allocs/op
BenchmarkBlake2b_256b-4                        4824     1137.15 μs/op       32 B/op   1 allocs/op
BenchmarkBlake2b_512b-4                        4567     1476.38 μs/op       64 B/op   1 allocs/op
BenchmarkSHA1_160b-4                           3822     1362.28 μs/op       24 B/op   1 allocs/op
BenchmarkMD5_128b-4                            3396     1807.97 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_512b-4                            2646     2254.14 μs/op       64 B/op   1 allocs/op
BenchmarkFnv1a_128b-4                          1767     3532.62 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_256b-4                            1672     3316.44 μs/op       32 B/op   1 allocs/op
BenchmarkFnv1_128b-4                           1658     3659.57 μs/op       16 B/op   1 allocs/op
BenchmarkSHA3_256b-4                           1396     3958.20 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_512b-4                            783     7036.49 μs/op      576 B/op   3 allocs/op
BenchmarkWhirlpool_512b-4                       178    35308.73 μs/op       64 B/op   1 allocs/op
ok  	github.com/KEINOS/go-blake3-example	225.237s
```

```text
name                                μs/op
CRC_32b-4                            71.5 ± 0%
XxhashCespare_8byte_64bit-4          87.4 ± 0%
XxhashOneOfOne_8byte_64bit_Sum64-4    108 ± 0%
XxhashOneOfOne_8byte_64bit-4          101 ± 0%
XxhashCespare_8byte_64bit_Sum64-4     105 ± 0%
XxhashOneOfOne_4byte_32bit_Sum32-4    204 ± 0%
MMH3_128b-4                           253 ± 0%
XxhashOneOfOne_4byte_32bit-4          292 ± 0%
Blake3_verZeebo_256b-4                414 ± 0%
Blake3_verZeebo_256b_Sum256-4         488 ± 0%
Blake3_verZeebo_512b-4                588 ± 0%
Blake3_verLuke_256b_Sum256-4          597 ± 0%
Blake3_verLuke_512b_Sum512-4          676 ± 0%
Blake3_verZeebo_512b_Sum512-4         562 ± 0%
Blake3_verLuke_512b-4                 727 ± 0%
Blake3_verLuke_256b-4                 637 ± 0%
Adler_32b-4                           711 ± 0%
Blake2b_256b-4                      1.14k ± 0%
Blake2b_512b-4                      1.48k ± 0%
SHA1_160b-4                         1.36k ± 0%
MD5_128b-4                          1.81k ± 0%
SHA_512b-4                          2.25k ± 0%
Fnv1a_128b-4                        3.53k ± 0%
SHA_256b-4                          3.32k ± 0%
Fnv1_128b-4                         3.66k ± 0%
SHA3_256b-4                         3.96k ± 0%
SHA3_512b-4                         7.04k ± 0%
Whirlpool_512b-4                    35.3k ± 0%

name                                alloc/op
CRC_32b-4                           8.00B ± 0%
XxhashCespare_8byte_64bit-4         8.00B ± 0%
XxhashOneOfOne_8byte_64bit_Sum64-4  0.00B
XxhashOneOfOne_8byte_64bit-4        8.00B ± 0%
XxhashCespare_8byte_64bit_Sum64-4   0.00B
XxhashOneOfOne_4byte_32bit_Sum32-4  0.00B
MMH3_128b-4                         16.0B ± 0%
XxhashOneOfOne_4byte_32bit-4        8.00B ± 0%
Blake3_verZeebo_256b-4              64.0B ± 0%
Blake3_verZeebo_256b_Sum256-4       0.00B
Blake3_verZeebo_512b-4              0.00B
Blake3_verLuke_256b_Sum256-4        0.00B
Blake3_verLuke_512b_Sum512-4        0.00B
Blake3_verZeebo_512b_Sum512-4       0.00B
Blake3_verLuke_512b-4                512B ± 0%
Blake3_verLuke_256b-4                256B ± 0%
Adler_32b-4                         8.00B ± 0%
Blake2b_256b-4                      32.0B ± 0%
Blake2b_512b-4                      64.0B ± 0%
SHA1_160b-4                         24.0B ± 0%
MD5_128b-4                          16.0B ± 0%
SHA_512b-4                          64.0B ± 0%
Fnv1a_128b-4                        16.0B ± 0%
SHA_256b-4                          32.0B ± 0%
Fnv1_128b-4                         16.0B ± 0%
SHA3_256b-4                          512B ± 0%
SHA3_512b-4                          576B ± 0%
Whirlpool_512b-4                    64.0B ± 0%

name                                allocs/op
CRC_32b-4                            1.00 ± 0%
XxhashCespare_8byte_64bit-4          1.00 ± 0%
XxhashOneOfOne_8byte_64bit_Sum64-4   0.00
XxhashOneOfOne_8byte_64bit-4         1.00 ± 0%
XxhashCespare_8byte_64bit_Sum64-4    0.00
XxhashOneOfOne_4byte_32bit_Sum32-4   0.00
MMH3_128b-4                          1.00 ± 0%
XxhashOneOfOne_4byte_32bit-4         1.00 ± 0%
Blake3_verZeebo_256b-4               2.00 ± 0%
Blake3_verZeebo_256b_Sum256-4        0.00
Blake3_verZeebo_512b-4               0.00
Blake3_verLuke_256b_Sum256-4         0.00
Blake3_verLuke_512b_Sum512-4         0.00
Blake3_verZeebo_512b_Sum512-4        0.00
Blake3_verLuke_512b-4                1.00 ± 0%
Blake3_verLuke_256b-4                1.00 ± 0%
Adler_32b-4                          1.00 ± 0%
Blake2b_256b-4                       1.00 ± 0%
Blake2b_512b-4                       1.00 ± 0%
SHA1_160b-4                          1.00 ± 0%
MD5_128b-4                           1.00 ± 0%
SHA_512b-4                           1.00 ± 0%
Fnv1a_128b-4                         1.00 ± 0%
SHA_256b-4                           1.00 ± 0%
Fnv1_128b-4                          1.00 ± 0%
SHA3_256b-4                          3.00 ± 0%
SHA3_512b-4                          3.00 ± 0%
Whirlpool_512b-4                     1.00 ± 0%
```
