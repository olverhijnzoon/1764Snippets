package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Crypto encodingBase64")

	data := "1764Snippets!\"§$%&/('Ä')"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}
