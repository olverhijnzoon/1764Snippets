package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	seq := make([]int, n)
	seq[0], seq[1] = 0, 1

	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	return seq
}

// This function transforms the Fibonacci sequence into a new sequence where each term is the reciprocal of the corresponding term in the Fibonacci sequence. This is a simple way to represent the concept of phase-matching, where each term in the sequence represents a different phase-matching condition.
func phaseMatching(n int) []float64 {
	fibSeq := fibonacci(n)
	phaseSeq := make([]float64, n)

	for i := range fibSeq {
		if fibSeq[i] != 0 {
			phaseSeq[i] = 1 / float64(fibSeq[i])
		} else {
			phaseSeq[i] = 0
		}
	}

	return phaseSeq
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Fibonacci")
	fmt.Println(fibonacci(10))
	fmt.Println(phaseMatching(25))
}
