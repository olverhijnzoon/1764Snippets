package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto SHA256")

	/*
		This snippet explores the crypto/sha256 package from the Go standard library, which provides an implementation of the SHA-256 (Secure Hash Algorithm) hash algorithm.

		Hashes can prove the integrity of data by comparing a provided hash from the source with a self-generated hash of the data. This works because it is A) unlikely to find two different input values that result in the same hash output, making it "collision-resistant," and B) infeasible to generate the input given only the hash provided by the source, making it "one-way." This distinguishes hashing from encryption and encoding.
	*/

	// Define a string message
	message := "1764"

	// Convert the message to bytes
	messageBytes := []byte(message)

	// Compute the SHA-256 hash
	hash := sha256.Sum256(messageBytes)

	// Print the resulting hash in hexadecimal format
	fmt.Printf("SHA-256 hash of \"1764\": %x\n", hash)

	// Provided SHA-256 hash
	providedHash, err := hex.DecodeString("9f273a349b224b830bf816e38baa3d1ffd849abc889d57fbba8024042e5693db")
	if err != nil {
		panic(err)
	}

	// Compare the computed hash with the provided hash
	isEqual := bytes.Equal(hash[:], providedHash)

	// Print true if the two hashes are equal and false if they are not
	fmt.Print("Are provided and calculated hashes equal? ")
	fmt.Println(isEqual)
}
