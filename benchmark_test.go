package main

import (
	"hash/fnv"
	"testing"

	"golang.org/x/crypto/sha3"
	"lukechampine.com/blake3"
)

// inputData holds 1MB(1e6) size of data created by testData() function.
var inputData []byte

// ============================================================================
//  Tests
// ============================================================================

// ----------------------------------------------------------------------------
//  FNV
// ----------------------------------------------------------------------------

func BenchmarkFNV1_4Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := fnv.New32()
		_, _ = h.Write(data)
		_ = h.Sum(nil)
	}
}

func BenchmarkFNV1A_4Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := fnv.New32a()
		_, _ = h.Write(data)
		_ = h.Sum(nil)
	}
}

// ----------------------------------------------------------------------------
//  BLAKE3
// ----------------------------------------------------------------------------

func BenchmarkBlake3_32Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blake3.Sum256(data)
	}
}

func BenchmarkBlake3_64Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blake3.Sum512(data)
	}
}

// ----------------------------------------------------------------------------
//  SHA3
// ----------------------------------------------------------------------------

func BenchmarkSHA3_32Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sha3.Sum256([]byte(data))
	}
}

func BenchmarkSHA3_64Byte(b *testing.B) {
	data := testData(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sha3.Sum512([]byte(data))
	}
}

// ============================================================================
//  Helper Functions
// ============================================================================

// The testData creates 1000,000 bytes= 1MB(1e6) size of data.
func testData(b *testing.B) []byte {
	b.Helper()

	if len(inputData) == 0 {
		inputData = make([]byte, 1e6)

		for i := range inputData {
			inputData[i] = byte(i % 251)
		}
	}

	return inputData
}
