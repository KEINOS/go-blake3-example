package main

import (
	"fmt"
	"hash/fnv"
	"io"
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

// ----------------------------------------------------------------------------
//  BLAKE3 (Luke Champine ver.)
// ----------------------------------------------------------------------------

func TestBlake3Luke_32byte_256bit_Sum256(t *testing.T) {
	input := "This is a string"

	valByte := blake3Luke.Sum256([]byte(input))

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Luke_32byte_256bit(t *testing.T) {
	input := "This is a string"

	h := blake3Luke.New(256, nil)

	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write data. Err: %v", err)
	}

	valByte := h.Sum(nil)[:32]

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Luke_64byte_512bit(t *testing.T) {
	input := "This is a string"

	hashByte := blake3Luke.Sum512([]byte(input))

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44"
	actual := fmt.Sprintf("%x", hashByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Luke_64byte_512bit_Sum512(t *testing.T) {
	input := "This is a string"

	h := blake3Luke.New(512, nil)

	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write data. Err: %v", err)
	}

	valByte := h.Sum(nil)[:64]

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

// ----------------------------------------------------------------------------
//  BLAKE3 (Zeebo ver.)
// ----------------------------------------------------------------------------

func TestBlake3Zeebo_32byte_256bit_Sum256(t *testing.T) {
	input := "This is a string"

	valByte := blake3Zeebo.Sum256([]byte(input))

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Zeebo_32byte_256bit(t *testing.T) {
	input := "This is a string"

	h := blake3Zeebo.New()

	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write data. Err: %v", err)
	}

	valByte := h.Sum(nil)[:32]

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Zeebo_64byte_512bit(t *testing.T) {
	input := "This is a string"

	hashByte := blake3Zeebo.Sum512([]byte(input))

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44"
	actual := fmt.Sprintf("%x", hashByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
}

func TestBlake3Zeebo_64byte_512bit_Sum512(t *testing.T) {
	input := "This is a string"

	h := blake3Zeebo.New()

	if _, err := h.Write([]byte(input)); err != nil {
		t.Fatalf("failed to write data. Err: %v", err)
	}

	d := h.Digest()

	out := make([]byte, 64)
	d.Seek(0, io.SeekStart)
	_, _ = d.Read(out)

	valByte := out

	// Expect value was taken from R implementation: https://github.com/dirkschumacher/blake3
	expect := "718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44"
	actual := fmt.Sprintf("%x", valByte)

	assert.Equal(t, expect, actual, "wrong hash value returned")
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
