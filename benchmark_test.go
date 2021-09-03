package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"io"
	"testing"

	"github.com/jzelinskie/whirlpool"
	"github.com/reusee/mmh3"
	blake3Zeebo "github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
	blake3Luke "lukechampine.com/blake3"
)

// inputData holds 1MB(1e6) size of data created by testData() function.
var inputData []byte

// ============================================================================
//  Tests
// ============================================================================

func BenchmarkAdler_32b(b *testing.B) {
	benchHashWriteAndSum(b, adler32.New())
}

func BenchmarkBlake2b_256b(b *testing.B) {
	h, err := blake2b.New256(nil)
	if err != nil {
		b.Fatal(err)
	}
	benchHashWriteAndSum(b, h)
}

func BenchmarkBlake2b_512b(b *testing.B) {
	h, err := blake2b.New512(nil)
	if err != nil {
		b.Fatal(err)
	}
	benchHashWriteAndSum(b, h)
}

func BenchmarkBlake3_verLuke_256b(b *testing.B) {
	h := blake3Luke.New(256, nil)
	benchHashWriteAndSum(b, h)
}

func BenchmarkBlake3_verLuke_256b_Sum256(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		blake3Luke.Sum256(data)
	}
}

func BenchmarkBlake3_verLuke_512b(b *testing.B) {
	h := blake3Luke.New(512, nil)
	benchHashWriteAndSum(b, h)
}

func BenchmarkBlake3_verLuke_512b_Sum512(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		blake3Luke.Sum512(data)
	}
}

func BenchmarkBlake3_verZeebo_256b(b *testing.B) {
	benchHashWriteAndSum(b, blake3Zeebo.New())
}

func BenchmarkBlake3_verZeebo_256b_Sum256(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		blake3Zeebo.Sum256(data)
	}
}

func BenchmarkBlake3_verZeebo_512b(b *testing.B) {
	data := testData(b)

	h := blake3Zeebo.New()

	b.ResetTimer()

	// To create longer or shorter than 32 byte zeebo needs
	// to use Digest
	for n := 0; n < b.N; n++ {
		if _, err := h.Write([]byte(data)); err != nil {
			b.Fatalf("failed to write data. Err: %v", err)
		}

		d := h.Digest()

		out := make([]byte, 64)
		d.Seek(0, io.SeekStart)
		_, _ = d.Read(out)
	}
}

func BenchmarkBlake3_verZeebo_512b_Sum512(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		blake3Zeebo.Sum512(data)
	}
}

func BenchmarkMMH3_128b(b *testing.B) {
	benchHashWriteAndSum(b, mmh3.New128())
}

func BenchmarkCRC_32b(b *testing.B) {
	benchHashWriteAndSum(b, crc32.NewIEEE())
}

func BenchmarkFnv1_128b(b *testing.B) {
	benchHashWriteAndSum(b, fnv.New128())
}

func BenchmarkFnv1a_128b(b *testing.B) {
	benchHashWriteAndSum(b, fnv.New128a())
}

func BenchmarkMD5_128b(b *testing.B) {
	benchHashWriteAndSum(b, md5.New())
}

func BenchmarkSHA1_160b(b *testing.B) {
	benchHashWriteAndSum(b, sha1.New())
}

func BenchmarkSHA_256b(b *testing.B) {
	benchHashWriteAndSum(b, sha256.New())
}

func BenchmarkSHA_512b(b *testing.B) {
	benchHashWriteAndSum(b, sha512.New())
}

func BenchmarkSHA3_256b(b *testing.B) {
	benchHashWriteAndSum(b, sha3.New256())
}

func BenchmarkSHA3_512b(b *testing.B) {
	benchHashWriteAndSum(b, sha3.New512())
}

func BenchmarkWhirlpool_512b(b *testing.B) {
	benchHashWriteAndSum(b, whirlpool.New())
}

// ============================================================================
//  Helper Functions
// ============================================================================

func benchHashWriteAndSum(b *testing.B, h hash.Hash) {
	data := testData(b)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, err := h.Write(data); err != nil {
			b.Fatalf("failed to write data. Err: %v", err)
		}
		_ = h.Sum(nil)
	}
}

// The testData creates 1,000,000 bytes= 1MB (1e6) size of data.
// The values are consistent and not random.
func testData(b *testing.B) []byte {
	b.Helper()

	// Initialize data or use initialized data
	if len(inputData) == 0 {
		inputData = make([]byte, 1e6)

		for i := range inputData {
			inputData[i] = byte(i % 251)
		}
	}

	return inputData
}
