package main

import "fmt"

const (
	A = iota
	C
	T
	G
)

// next returns the next base in the sequence
func next(base int) int {
	if base == G {
		return A // Wrap around to A when G is reached
	}
	return base + 1
}

// baseToString converts a base constant to its string representation
func baseToString(base int) string {
	switch base {
	case A:
		return "A"
	case C:
		return "C"
	case T:
		return "T"
	case G:
		return "G"
	default:
		return "Unknown"
	}
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Iterator")

	base := A
	for i := 0; i < 10; i++ {
		fmt.Println(baseToString(base))
		base = next(base)
	}
}
