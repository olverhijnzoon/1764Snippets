package main

import (
	"crypto/rand"
	"fmt"

	"filippo.io/edwards25519"
)

// randBytes returns a slice of n random bytes
func randBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto Edwards25519")

	/*
		This snippet explores the `edwards25519` package from `filippo.io`, which is an implementation of the twisted Edwards curve, a specific type of elliptic curve that is used in modern cryptography.

		Elliptic curves are used in cryptography due to their mathematical properties that allow for efficient computation and strong security. They are used in various cryptographic protocols, including key exchange, digital signatures, and encryption.

		The `edwards25519` curve is equivalent to `Curve25519`, a curve designed for high-speed high-security elliptic curve cryptography and is used by the `Ed25519` signature scheme. The `Ed25519` signature scheme is widely used due to its strong security properties and efficiency. The curve equation is:

			-x^2 + y^2 = 1 + -(121665/121666) * x^2 * y^2

		As of Go 1.17, it was merged back into the Go standard library. There are higher-level libraries like crypto/ed25519 for signatures or golang.org/x/crypto/curve25519 for Diffie-Hellman.

		In the edwards25519 package, points are represented as four coordinates (X, Y, Z, T), and the actual point on the curve is the ratio (X/Z, Y/Z). The identity point is usually represented as (0, 1, 1, 0).

		In this code, we start by creating two identity points on the curve. We then add these two identity points together using the Add function. The result of this operation should be another identity point. To verify this, we use the Equal function to check if the resulting point from the addition is equal to one of the original identity points. Since we are adding two identity points, we expect the result to be equal to the identity point, so Equal should return 1.

		Next, we generate a random scalar. In elliptic curve cryptography, a scalar is a number used to multiply a point on the curve. This is done using the NewScalar and SetUniformBytes functions. The SetUniformBytes function expects a 64-byte (512-bit) uniformly random input, which is then hashed and reduced modulo the group order to produce a scalar.

		We then multiply the base point of the curve by this scalar using the ScalarBaseMult function. This operation is fundamental in many cryptographic protocols, as it allows us to generate a new point on the curve that can only be derived from the original point by someone who knows the scalar.

		Finally, we use the Equal function again to check if the resulting point from the scalar multiplication is equal to the identity point. Unless the scalar is zero, this should not be the case, so we expect Equal to return 0.

		This is a slightly modified version of the original 42Snippets golangedwards25519 snippet "Golang Edwards25519".
	*/

	// Create two new identity points
	identityPoint1 := edwards25519.NewIdentityPoint()
	identityPoint2 := edwards25519.NewIdentityPoint()

	// Print the resulting points
	fmt.Println(identityPoint1)
	fmt.Println(identityPoint2)

	// Add the two identity points
	sum := new(edwards25519.Point).Add(identityPoint1, identityPoint2)

	// When you print a point using fmt.Println, it prints the internal representation of the point, which includes all four coordinates. The Add function may result in a point that is equivalent to the identity point, but with different internal coordinates due to the way the addition is performed.
	fmt.Println(sum)

	// Check if the sum is equal to the identity point
	isEqual := sum.Equal(identityPoint1)

	// Print 1 if the two points are equal and 0 if they are not
	fmt.Println(isEqual)

	// Generate a random scalar
	scalar, err := edwards25519.NewScalar().SetUniformBytes(randBytes(64))
	if err != nil {
		panic(err)
	}

	// Print the resulting point
	fmt.Println(scalar)

	// Multiply the base point by the scalar
	anyPoint := new(edwards25519.Point).ScalarBaseMult(scalar)

	// Print the resulting point
	fmt.Println(anyPoint)

	// Check if the sum is equal to the identity point
	isEqual0 := anyPoint.Equal(identityPoint1)

	// Should print 0 unless it has somehow randomly generated a zero scalar
	fmt.Println(isEqual0)
}
