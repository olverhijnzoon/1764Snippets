package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto encodingBase64")

	/*
		One topic is Padding, the "=" character at the end of the encoded string is used when the number of bytes to encode is not divisible by 3. If there is only one remaining byte to encode, "==" padding characters are added at the end. If the number of bytes to encode is divisible by 3, no padding is done.

		Another topic is URL Safe Encoding as Base64 includes ‘+’ and ‘/’ characters which are not safe for URLs.
	*/

	// No padding but / at the end
	data := "1764Snippets!\"§$%&/('Ä'+?"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	// Padding =
	dataSinglePadding := "1764Snippets!\"§$%&/('Ä'+"
	encodedSinglePadding := base64.StdEncoding.EncodeToString([]byte(dataSinglePadding))
	fmt.Println("Encoded:", encodedSinglePadding)

	decodedSinglePadding, err := base64.StdEncoding.DecodeString(encodedSinglePadding)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Decoded:", string(decodedSinglePadding))

	// Padding ==
	dataDoublePadding := "1764Snippets!\"§$%&/('Ä'"
	encodedDoublePadding := base64.StdEncoding.EncodeToString([]byte(dataDoublePadding))
	fmt.Println("Encoded:", encodedDoublePadding)

	decodedDoublePadding, err := base64.StdEncoding.DecodeString(encodedDoublePadding)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Decoded:", string(decodedDoublePadding))

	// URL Safe Encoding + /
	urlSafeEncoded := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)

	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("URL Safe Decode error:", err)
		return
	}
	fmt.Println("URL Safe Decoded:", string(urlSafeDecoded))
}
