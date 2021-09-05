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
    - We compared the below BLAKE3 implementations in Go (Unofficial)
    - Luke Champine's BLAKE3 for Go
        - `go get lukechampine.com/blake3`
        - [Repo](https://github.com/lukechampine/blake3) | lukechampine @ GitHub
        - [Document](https://pkg.go.dev/lukechampine.com/blake3) @ pkg.go.dev
    - Zeebo's BLAKE3 for Go
        - `go get github.com/zeebo/blake3`
        - [Repo](https://github.com/zeebo/blake3) | zeebo @ GitHub
        - [Document](https://pkg.go.dev/github.com/zeebo/blake3) @ pkg.go.dev

## Bench Results

OMG ... even `blake3.Sum512()`(64Bytes long hash) is faster than `fnv.New32()`(4Bytes long hash, [FNV1](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function)).

Blake3-512bit (Luke ver) was 45,450 iterations/10s max and FNV1a-128 was 5,086 iterations/10s max ... Am I [doing something](./benchmark_test.go) wrong?

```shellsession
$ go test -failfast -benchmem -shuffle=on -benchtime=5s -count=5 -bench . | go-prettybench -sort iter
-test.shuffle 1630647622685374698
goos: linux
goarch: amd64
pkg: github.com/KEINOS/go-blake3-example
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
PASS
benchmark                                      iter         time/iter   bytes alloc        allocs
---------                                      ----         ---------   -----------        ------
BenchmarkCRC_32b-2                            90219       66.24 μs/op        8 B/op   1 allocs/op
BenchmarkCRC_32b-2                            89403       66.56 μs/op        8 B/op   1 allocs/op
BenchmarkCRC_32b-2                            88344       66.93 μs/op        8 B/op   1 allocs/op
BenchmarkCRC_32b-2                            86281       69.86 μs/op        8 B/op   1 allocs/op
BenchmarkCRC_32b-2                            83502       72.33 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-2    69592       86.39 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-2    69496       87.76 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit-2          69427       86.92 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-2   69302       87.54 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-2   68941       86.82 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit-2          68936       87.50 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-2   68673       86.42 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit-2          68493       86.66 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-2    68037       86.12 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-2   67582       87.12 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit_Sum64-2   67116       91.52 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-2    66859       93.44 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashCespare_8byte_64bit-2          66841       91.24 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit-2          64641       89.57 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashCespare_8byte_64bit_Sum64-2    61558       98.97 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-2         61521       99.56 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-2         61393       98.27 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-2         60975       97.18 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-2         60439      101.53 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_8byte_64bit-2         57440      104.51 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-2   31035      193.81 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-2   30363      195.32 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-2   29811      193.89 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-2   27054      194.59 μs/op        0 B/op   0 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit_Sum32-2   26088      208.95 μs/op        0 B/op   0 allocs/op
BenchmarkMMH3_128b-2                          25233      236.69 μs/op       16 B/op   1 allocs/op
BenchmarkMMH3_128b-2                          25095      238.95 μs/op       16 B/op   1 allocs/op
BenchmarkMMH3_128b-2                          24853      242.64 μs/op       16 B/op   1 allocs/op
BenchmarkMMH3_128b-2                          24843      240.52 μs/op       16 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-2         22977      259.93 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-2         22918      259.45 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-2         22827      267.43 μs/op        8 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-2         22725      261.07 μs/op        8 B/op   1 allocs/op
BenchmarkMMH3_128b-2                          22676      239.52 μs/op       16 B/op   1 allocs/op
BenchmarkXxhashOneOfOne_4byte_32bit-2         20400      264.64 μs/op        8 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b-2               15711      381.57 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b-2               15559      385.23 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b-2               14119      385.77 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b-2               13540      429.84 μs/op       64 B/op   2 allocs/op
BenchmarkBlake3_verZeebo_256b-2               12685      419.09 μs/op       64 B/op   2 allocs/op
BenchmarkAdler_32b-2                          12655      473.27 μs/op        8 B/op   1 allocs/op
BenchmarkAdler_32b-2                          12637      475.49 μs/op        8 B/op   1 allocs/op
BenchmarkAdler_32b-2                          12624      490.90 μs/op        8 B/op   1 allocs/op
BenchmarkAdler_32b-2                          12590      474.22 μs/op        8 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-2        12007      497.91 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-2        11367      499.17 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-2        10874      511.25 μs/op        0 B/op   0 allocs/op
BenchmarkAdler_32b-2                          10870      499.71 μs/op        8 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b-2                10789      530.54 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-2         10758      529.80 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-2         10710      531.63 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-2         10615      553.61 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-2                10609      534.17 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-2        10593      497.49 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-2         10573      528.68 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-2               10510      521.88 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-2               10490      538.74 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b-2                10000      535.44 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verLuke_512b-2                10000      536.30 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-2         10000      528.61 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-2                10000      532.66 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-2         10000      535.04 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-2                10000      540.10 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verLuke_256b-2                10000      530.26 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-2        10000      501.89 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-2        10000      587.27 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b-2                10000      531.65 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_512b-2               10000      527.25 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-2         10000      528.59 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-2        10000      504.74 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-2               10000      525.22 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-2         10000      528.68 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b-2                10000      536.27 μs/op      256 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-2        10000      501.58 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b_Sum512-2        10000      501.33 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b-2                10000      534.49 μs/op      512 B/op   1 allocs/op
BenchmarkBlake3_verZeebo_256b_Sum256-2        10000      509.91 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_256b_Sum256-2         10000      525.02 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verLuke_512b_Sum512-2          9908      529.79 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_verZeebo_512b-2                9622      653.41 μs/op        0 B/op   0 allocs/op
BenchmarkBlake2b_256b-2                        5392     1125.82 μs/op       32 B/op   1 allocs/op
BenchmarkBlake2b_512b-2                        5383     1135.67 μs/op       64 B/op   1 allocs/op
BenchmarkBlake2b_512b-2                        5343     1123.82 μs/op       64 B/op   1 allocs/op
BenchmarkBlake2b_512b-2                        5326     1144.87 μs/op       64 B/op   1 allocs/op
BenchmarkBlake2b_512b-2                        5142     1119.85 μs/op       64 B/op   1 allocs/op
BenchmarkBlake2b_512b-2                        5112     1138.34 μs/op       64 B/op   1 allocs/op
BenchmarkBlake2b_256b-2                        5106     1332.33 μs/op       32 B/op   1 allocs/op
BenchmarkSHA1_160b-2                           5083     1189.95 μs/op       24 B/op   1 allocs/op
BenchmarkSHA1_160b-2                           5080     1186.62 μs/op       24 B/op   1 allocs/op
BenchmarkSHA1_160b-2                           4915     1196.98 μs/op       24 B/op   1 allocs/op
BenchmarkBlake2b_256b-2                        4831     1240.18 μs/op       32 B/op   1 allocs/op
BenchmarkSHA1_160b-2                           4768     1198.96 μs/op       24 B/op   1 allocs/op
BenchmarkBlake2b_256b-2                        4640     1152.73 μs/op       32 B/op   1 allocs/op
BenchmarkBlake2b_256b-2                        4590     1158.44 μs/op       32 B/op   1 allocs/op
BenchmarkSHA1_160b-2                           4453     1297.59 μs/op       24 B/op   1 allocs/op
BenchmarkMD5_128b-2                            3578     1708.95 μs/op       16 B/op   1 allocs/op
BenchmarkMD5_128b-2                            3570     1697.82 μs/op       16 B/op   1 allocs/op
BenchmarkMD5_128b-2                            3558     1680.34 μs/op       16 B/op   1 allocs/op
BenchmarkMD5_128b-2                            3212     1833.90 μs/op       16 B/op   1 allocs/op
BenchmarkMD5_128b-2                            3094     1752.61 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_512b-2                            2835     2134.08 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_512b-2                            2814     2146.41 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_512b-2                            2637     2337.40 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_512b-2                            2472     2216.42 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_512b-2                            2334     2220.89 μs/op       64 B/op   1 allocs/op
BenchmarkSHA_256b-2                            1989     3025.40 μs/op       32 B/op   1 allocs/op
BenchmarkSHA_256b-2                            1983     3039.82 μs/op       32 B/op   1 allocs/op
BenchmarkSHA_256b-2                            1981     3036.94 μs/op       32 B/op   1 allocs/op
BenchmarkSHA_256b-2                            1774     3142.80 μs/op       32 B/op   1 allocs/op
BenchmarkFnv1a_128b-2                          1755     3424.04 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1a_128b-2                          1726     3469.33 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1a_128b-2                          1724     3435.15 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1a_128b-2                          1718     3440.90 μs/op       16 B/op   1 allocs/op
BenchmarkSHA_256b-2                            1718     3785.94 μs/op       32 B/op   1 allocs/op
BenchmarkFnv1_128b-2                           1708     3510.60 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1_128b-2                           1705     3543.74 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1_128b-2                           1704     3544.08 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1_128b-2                           1666     3529.24 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1a_128b-2                          1662     3602.80 μs/op       16 B/op   1 allocs/op
BenchmarkFnv1_128b-2                           1644     3610.79 μs/op       16 B/op   1 allocs/op
BenchmarkSHA3_256b-2                           1610     3675.46 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_256b-2                           1610     3689.29 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_256b-2                           1606     3683.08 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_256b-2                           1508     3846.73 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_256b-2                           1484     3835.98 μs/op      512 B/op   3 allocs/op
BenchmarkSHA3_512b-2                            877     6897.34 μs/op      576 B/op   3 allocs/op
BenchmarkSHA3_512b-2                            864     6926.91 μs/op      576 B/op   3 allocs/op
BenchmarkSHA3_512b-2                            847     6982.35 μs/op      576 B/op   3 allocs/op
BenchmarkSHA3_512b-2                            843     7193.62 μs/op      576 B/op   3 allocs/op
BenchmarkSHA3_512b-2                            817     7022.99 μs/op      576 B/op   3 allocs/op
BenchmarkWhirlpool_512b-2                       194    30886.98 μs/op       64 B/op   1 allocs/op
BenchmarkWhirlpool_512b-2                       193    31062.21 μs/op       64 B/op   1 allocs/op
BenchmarkWhirlpool_512b-2                       192    30965.45 μs/op       64 B/op   1 allocs/op
BenchmarkWhirlpool_512b-2                       174    31147.90 μs/op       64 B/op   1 allocs/op
BenchmarkWhirlpool_512b-2                       171    32797.90 μs/op       64 B/op   1 allocs/op
ok  github.com/KEINOS/go-blake3-example 228.463s
```

- This repo is GitHub Codespaces compatible. Try your own online on GitHub.
