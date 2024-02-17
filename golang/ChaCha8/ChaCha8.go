package main

import (
	"encoding/hex"
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto ChaCha8")

	/*
		This snippet looks into https://pkg.go.dev/math/rand/v2@go1.22.0 which introduces first “v2” package in the standard library. Starting from the release of Go 1.22, the top-level functions of the math/rand package, as well as the Go runtime system, utilize the ChaCha8 algorithm for generating random numbers, but only when these functions are not explicitly seeded. Meaning, if you don’t provide a specific seed, ChaCha8 is the default algorithm used for randomness in these contexts.

		This snippet is intentionally not in the crypto folder of this repository as math/rand/v2 pseudo-random number generators should not be used for cryptography. Documentation states clearly that this for simulation or similar.

		The ChaCha8 algorithm is a type of pseudorandom number generator (PRNG). It takes a fixed seed as input and generates a sequence of numbers that appear random. However, given the same seed, ChaCha8 will always produce the same sequence of numbers. This is a property of all PRNGs, not just ChaCha8.
	*/

	// Creates a fixed-seed for the ChaCha8 generator (dont do it like this in real code)
	var seed [32]byte
	for i := range seed {
		seed[i] = byte(i)
		// seed[i] = byte(i) ^ byte(i) XOR it for 0000
	}

	// Print the seed. It should not change when its done like this.
	fmt.Println("Seed:", hex.EncodeToString(seed[:]))

	chacha8Generator := rand.NewChaCha8(seed)
	randomGenerator := rand.New(chacha8Generator)

	// A fixed seed results in same output on every run
	fmt.Println("Random int:", randomGenerator.Int())
	fmt.Println("Random float64:", randomGenerator.Float64())
	fmt.Println("Random int32 in range [0,10):", randomGenerator.Int32N(10))
}
