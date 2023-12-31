package main

import (
	"fmt"
	"math"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialStirlingApproximation(n float64) float64 {
	if n < 10 {
		fmt.Println("Warning: For small numbers, it's better to use the normal factorial function instead of Stirling's approximation.")
	}
	return math.Sqrt(2*math.Pi*n) * math.Pow(n/math.E, n)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combinationWithRepetition(k, r int) int {
	return factorial(k) / (factorial(r) * factorial(k-r))
}

func combinationWithRepetitionStirling(k, r float64) float64 {
	// Could be more beatiful
	return factorialStirlingApproximation(k) / (factorialStirlingApproximation(r) * factorialStirlingApproximation(k-r))
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Combination with Repetition")

	fmt.Println("Factorial of 5 is", factorial(5))
	fmt.Println("Stirling's approximation of factorial for 5 is", factorialStirlingApproximation(5))
	fmt.Println("Combination of 5 and 3 is", combination(5, 3))
	fmt.Println("Permutation of 5 and 3 is", permutation(5, 3))
	fmt.Println("Combination with repetition of 5 and 3 is", combinationWithRepetition(5, 3))
	fmt.Println("Combination with repetition of 20 and 10 is", combinationWithRepetition(20, 10))
	fmt.Println("Stirling's approximation of combination with repetition of 20 and 10 is", combinationWithRepetitionStirling(20, 10))
}
