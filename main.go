/*
Package main is an example usge of BLAKE3 in GO.

- [BLAKE3](https://en.wikipedia.org/wiki/BLAKE_(hash_function)#BLAKE3) implementation for Go by [@lukechampine](https://github.com/lukechampine/blake3).
*/
package main

import (
	"fmt"

	"lukechampine.com/blake3"
)

func main() {
	value := "This is a string"

	{
		valByte := blake3.Sum256([]byte(value))
		fmt.Println(
			"Expect:",
			"718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2",
		)
		fmt.Println("Actual:", fmt.Sprintf("%x", valByte))
	}

	{
		valByte := blake3.Sum512([]byte(value))
		fmt.Println(
			"Expect:",
			"718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f243a9903464658159cf8b216e79006e12ef3568851423fa7c97002cbb9ca4dc44",
		)
		fmt.Println("Actual:", fmt.Sprintf("%x", valByte))
	}

}
