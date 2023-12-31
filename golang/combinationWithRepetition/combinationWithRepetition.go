package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combinationWithRepetition(n, r int) int {
	return factorial(n+2*r-2) / (factorial(r) * factorial(n+r-2))
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Combination with Repetition")

	fmt.Println("Factorial of 5 is", factorial(5))
	fmt.Println("Combination of 5 and 3 is", combination(5, 3))
	fmt.Println("Permutation of 5 and 3 is", permutation(5, 3))
	fmt.Println("Combination with repetition of 5 and 3 is", combinationWithRepetition(5, 3))
}
