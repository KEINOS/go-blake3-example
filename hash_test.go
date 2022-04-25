package main

// Tests of hash functionality before benchmarking.

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"testing"

	xxHashOneOfOne "github.com/OneOfOne/xxhash"
	xxHashCespare "github.com/cespare/xxhash"
	"github.com/stretchr/testify/assert"
	blake3Zeebo "github.com/zeebo/blake3"
	blake3Luke "lukechampine.com/blake3"
)

// ----------------------------------------------------------------------------
//  FNV
// ----------------------------------------------------------------------------

// FNV1-32
func TestFNV1_4byte_32bit(t *testing.T) {
	input := "This is a string"

	h := fnv.New32()
	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write byte data")
	}
	valByte := h.Sum(nil)

	expect := "cc9ecaed"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual)
}

// FNV1a-32
func TestFNV1a_4byte_32bit(t *testing.T) {
	input := "This is a string"

	h := fnv.New32a()
	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write byte data")
	}
	valByte := h.Sum(nil)

	expect := "5eb7a185"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual)
}

// FNV1-64
func TestFNV1_8byte_64bit(t *testing.T) {
	input := "This is a string"

	h := fnv.New64()
	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write byte data")
	}
	valByte := h.Sum(nil)

	expect := "ed6b34f0f1866b4d"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual)
}

// FNV1a-64
func TestFNV1a_8byte_64bit(t *testing.T) {
	input := "This is a string"

	h := fnv.New64a()
	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write byte data")
	}
	valByte := h.Sum(nil)

	expect := "c8a0d501ca41c0a5"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual)
}

// ============================================================================
//  BLAKE3
// ============================================================================

// Helper function - It returns the golden values for BLAKE3 created from the
// official Rust implementation. See ./testdata directory for details.
func getTestCaseBlake3(t *testing.T, lenByte int) map[string]string {
	t.Helper()

	pathData := "./testdata/golden_data32.json"
	if lenByte == 64 {
		pathData = "./testdata/golden_data64.json"
	}

	dataRaw, err := os.ReadFile(pathData)
	if err != nil {
		t.Fatal(err)
	}

	// Key = intput, Value = expected hash value
	dataJSON := map[string]string{}

	err = json.Unmarshal(dataRaw, &dataJSON)
	if err != nil {
		t.Fatal(err)
	}

	return dataJSON
}

// ----------------------------------------------------------------------------
//  BLAKE3 (Luke Champine ver.)
// ----------------------------------------------------------------------------

func TestBlake3Luke_32byte_256bit_Sum256(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 32)

	for input, expect := range dataJSON {
		valByte := blake3Luke.Sum256([]byte(input))
		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Luke_32byte_256bit_write(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 32)

	for input, expect := range dataJSON {
		h := blake3Luke.New(256, nil)

		if _, err := h.Write([]byte(input)); err != nil {
			t.Fatalf("failed to write data. Err: %v", err)
		}

		valByte := h.Sum(nil)[:32]
		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Luke_64byte_512bit_Sum512(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 64)

	for input, expect := range dataJSON {
		hashByte := blake3Luke.Sum512([]byte(input))
		actual := fmt.Sprintf("%x", hashByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Luke_64byte_512bit_write(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 64)

	for input, expect := range dataJSON {
		h := blake3Luke.New(512, nil)

		if _, err := h.Write([]byte(input)); err != nil {
			t.Fatalf("failed to write data. Err: %v", err)
		}

		valByte := h.Sum(nil)[:64]
		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

// ----------------------------------------------------------------------------
//  BLAKE3 (Zeebo ver.)
// ----------------------------------------------------------------------------

func TestBlake3Zeebo_32byte_256bit_Sum256(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 32)

	for input, expect := range dataJSON {
		valByte := blake3Zeebo.Sum256([]byte(input))
		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Zeebo_32byte_256bit_Write(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 32)

	for input, expect := range dataJSON {
		h := blake3Zeebo.New()

		if _, err := h.Write([]byte(input)); err != nil {
			t.Fatalf("failed to write data. Err: %v", err)
		}

		valByte := h.Sum(nil)[:32]
		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Zeebo_64byte_512bit_Sum512(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 64)

	for input, expect := range dataJSON {
		hashByte := blake3Zeebo.Sum512([]byte(input))
		actual := fmt.Sprintf("%x", hashByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

func TestBlake3Zeebo_64byte_512bit_Write(t *testing.T) {
	dataJSON := getTestCaseBlake3(t, 64)

	for input, expect := range dataJSON {
		h := blake3Zeebo.New()

		if _, err := h.Write([]byte(input)); err != nil {
			t.Fatalf("failed to write data. Err: %v", err)
		}

		d := h.Digest()

		valByte := make([]byte, 64)
		d.Seek(0, io.SeekStart)
		_, _ = d.Read(valByte)

		actual := fmt.Sprintf("%x", valByte)

		assert.Equal(t, expect, actual, "wrong hash value returned")
	}
}

// ----------------------------------------------------------------------------
//  xxHash
// ----------------------------------------------------------------------------

func TestXxhashOneOfOne_8byte_64bit_Sum64(t *testing.T) {
	input := "This is a string"

	hashByte := xxHashOneOfOne.ChecksumString64(input)

	// Expect value was taken from online: https://asecuritysite.com/encryption/xxhash
	expect := fmt.Sprintf("%x", 0x0717e8ee90118ae1)
	actual := fmt.Sprintf("%x", hashByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestXxhashCespare_8byte_64bit_Sum64(t *testing.T) {
	input := "This is a string"

	hashByte := xxHashCespare.Sum64String(input)

	// Expect value was taken from online: https://asecuritysite.com/encryption/xxhash
	expect := fmt.Sprintf("%x", 0x0717e8ee90118ae1)
	actual := fmt.Sprintf("%x", hashByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}
