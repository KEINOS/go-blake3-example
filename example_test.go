package main

// ============================================================================
//  Example usage of BLAKE3 algorithm hash functions
// ============================================================================

import (
	"fmt"
	"io"
	"log"

	blake3Zeebo "github.com/zeebo/blake3"
	blake3Luke "lukechampine.com/blake3"
)

// ----------------------------------------------------------------------------
//   Example usage of "lukechampine.com/blake3" package
// ----------------------------------------------------------------------------

func Example_blake3Luke_Sum() {
	// import blake3Luke "lukechampine.com/blake3"

	input := "This is a string"

	// Example of Sum256 (32 bytes hash)
	// Sum256 returns the unkeyed BLAKE3 hash of b, truncated to 256 bits.
	valByte32 := blake3Luke.Sum256([]byte(input))
	fmt.Printf("%x\n", valByte32)

	// Example of Sum512 (64 bytes hash)
	// Sum512 returns the unkeyed BLAKE3 hash of b, truncated to 512 bits.
	valByte64 := blake3Luke.Sum512([]byte(input))
	fmt.Printf("%x\n", valByte64)

	// Output:
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44
}

func Example_blake3Luke_write() {
	// import blake3Luke "lukechampine.com/blake3"

	input := "This is a string"

	// Get 32 bytes hash
	{
		// New returns a Hasher for the specified size and key. Here the size is 32 bytes.
		// If key is nil, the hash is unkeyed. Otherwise, len(key) must be 32.
		h := blake3Luke.New(32, nil)

		if _, err := h.Write([]byte(input)); err != nil {
			log.Fatalf("failed to write data. Err: %v", err)
		}

		valByte := h.Sum(nil)
		fmt.Printf("%x\n", valByte)
	}

	// Get 64 bytes hash
	{
		// New returns a Hasher for the specified size and key. Here the size is 64 bytes.
		// If key is nil, the hash is unkeyed. Otherwise, len(key) must be 32.
		h := blake3Luke.New(64, nil)

		if _, err := h.Write([]byte(input)); err != nil {
			log.Fatalf("failed to write data. Err: %v", err)
		}

		valByte := h.Sum(nil)
		fmt.Printf("%x\n", valByte)
	}

	// Get both 32 and 64 bytes hash
	{
		h := blake3Luke.New(64, nil)

		if _, err := h.Write([]byte(input)); err != nil {
			log.Fatalf("failed to write data. Err: %v", err)
		}

		valByte32 := h.Sum(nil)[:32]
		fmt.Printf("%x\n", valByte32)

		valByte64 := h.Sum(nil)[:64]
		fmt.Printf("%x\n", valByte64)
	}

	// Output:
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44
}

// ----------------------------------------------------------------------------
//   Example usage of github.com/zeebo/blake3 package
// ----------------------------------------------------------------------------

func Example_blake3Zeebo_Sum() {
	// import blake3Zeebo "github.com/zeebo/blake3"

	input := "This is a string"

	// Example of Sum512 (32 bytes hash)
	hashByte32 := blake3Zeebo.Sum256([]byte(input))
	fmt.Printf("%x\n", hashByte32)

	// Example of Sum512 (64 bytes hash)
	hashByte64 := blake3Zeebo.Sum512([]byte(input))
	fmt.Printf("%x\n", hashByte64)

	// Output:
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44
}

func Example_blake3Zeebo_write() {
	// import blake3Zeebo "github.com/zeebo/blake3"

	input := "This is a string"

	// Instanciate new hasher
	h := blake3Zeebo.New()

	// Write data to hasher.
	// Write implements part of the hash.Hash interface. It never returns an error.
	if _, err := h.Write([]byte(input)); err != nil {
		log.Fatalf("failed to write data. Err: %v", err)
	}

	// Get 32 bytes hash
	{
		// Sum implements part of the hash.Hash interface. It appends the digest
		// of the Hasher to the provided buffer and returns it.
		valByte := h.Sum(nil)[:32]

		fmt.Printf("%x\n", valByte)
	}

	// Get 64 bytes hash
	{
		// Digest takes a snapshot of the hash state and returns an object that
		// can be used to read and seek through 2^64 "bytes" of digest output.
		d := h.Digest()

		valByte := make([]byte, 64)

		d.Seek(0, io.SeekStart)
		if _, err := d.Read(valByte); err != nil {
			log.Fatalf("failed to read data. Err: %v", err)
		}

		fmt.Printf("%x\n", valByte)
	}

	// Output:
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2
	// 718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44
}
