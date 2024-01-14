package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

/*
The GenerateKey function takes as input a curve (in this case, elliptic.P224) and a random number generator (in this case, rand.Reader), and it returns a private key and an error value.

The curve is called "P-224" because it is defined over a prime field with a modulus of 224 bits. The P-224 curve was chosen for use in ECDSA as it has a relatively small modulus with its 224 bits, which makes it efficient, but it is still considered to be secure against attacks.

The math/rand package of the Go standard library provides a cryptographically secure pseudo-random number generator (CSPRNG) which is used as source of randomness here.
*/
func genPPKeys() (privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, err error) {
	privateKey, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	publicKey = &privateKey.PublicKey
	return privateKey, publicKey, nil
}

/*
The signature is a tuple of values (r, s) that is derived from a (internally hashed) message, a private key and a random number. The function first generates this random number (called a "nonce") using a CSPRNG that is initialized using a key that is derived from the private key and a source of entropy.
*/
func sign(privateKey *ecdsa.PrivateKey, message []byte) (r, s *big.Int, err error) {
	r, s, err = ecdsa.Sign(rand.Reader, privateKey, message)
	if err != nil {
		return nil, nil, err
	}
	return r, s, nil
}

/*
The ecdsa.Verify function takes a public key, a message, and an "r" and "s" value as arguments and returns a boolean value indicating whether the signature is valid.
*/
func verify(publicKey *ecdsa.PublicKey, message []byte, r, s *big.Int) (valid bool, err error) {
	valid = ecdsa.Verify(publicKey, message, r, s)
	return valid, nil
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto Signature")

	/*
		This is an example of using the Elliptic Curve Digital Signature Algorithm (ECDSA) to generate a private/public key pair, sign a message with the private key, and verify the message with the public key.

		The code defines a function genPPKeys that generates a private/public key pair using the ecdsa package's GenerateKey function, a function sign that signs a message with a private key using the ecdsa package's Sign function, and a function verify that verifies a signature with a public key using the ecdsa package's Verify function. The main function then generates a private/public key pair, signs a message with the private key, and verifies the signature with the public key.

		This is a slightly modified version of the "Signature" snippet from the "Golang Snippets" repo predating the 42Snippets.
	*/

	// genPPKeys generates a pair of private and public keys using the Elliptic Curve Digital Signature Algorithm (ECDSA)
	privateKey, publicKey, err := genPPKeys()
	if err != nil {
		fmt.Printf("error generating keys: %s\n", err)
		return
	}

	// Creating a message
	message := []byte("This is a message to be signed")

	// This code is calling the sign function, which takes as input a private key and a message and returns a signature, which is a tuple of values (r, s).
	r, s, err := sign(privateKey, message)
	if err != nil {
		fmt.Printf("error signing message: %s\n", err)
		return
	}

	// In theory, the signature is then appended to the message and sent to the recipient, who can use the public key of the sender to verify the authenticity of the signature.
	valid, err := verify(publicKey, message, r, s)
	if err != nil {
		fmt.Printf("error verifying signature: %s\n", err)
		return
	}
	fmt.Printf("signature verification result: %t\n", valid)

	// Testing it with an invalid message
	invalidMessage := []byte("This is a different message")
	valid, err = verify(publicKey, invalidMessage, r, s)
	if err != nil {
		fmt.Printf("error verifying signature: %s\n", err)
		return
	}
	fmt.Printf("signature verification result with invalid message: %t\n", valid)
}
