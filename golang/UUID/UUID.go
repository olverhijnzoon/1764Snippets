package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang UUID")

	// Generate a version 1 UUID
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("Failed to generate UUID v1: %s\n", err)
	} else {
		fmt.Printf("Generated UUID v1: %s\n", uuidV1)
	}

	// Generate a version 4 UUID
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID v4: %s\n", err)
	} else {
		fmt.Printf("Generated UUID v4: %s\n", uuidV4)
	}
}
