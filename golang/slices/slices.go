package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Slices")

	/*
		This snippet demonstrates the usage of slices library and its updates introduced in Golang 1.22 https://tip.golang.org/doc/go1.22#go/version.

		The Concat function is used to concatenate slice1 and slice2. The Delete function is used to remove an element from slice3, and it shows how elements beyond the new length are zeroed out. The Compact function is used on slice4 to remove empty string elements and demonstrate the zeroing out of unused array slots. The Insert function is used with an out-of-range index to demonstrate that it now panics in such cases.
	*/

	slice1 := []int{1, 7, 6}
	slice2 := []int{4, 4, 2}
	concatenatedSlice := slices.Concat(slice1, slice2)
	fmt.Println("Concatenated Slice:", concatenatedSlice)

	slice3 := []string{"a", "b", "c", "d", "e"}
	modifiedSlice := slices.Delete(slice3, 2, 3)
	fmt.Println("Modified Slice after Delete:", modifiedSlice)
	fmt.Printf("Underlying array after Delete: %q\n", slice3)

	slice4 := []string{"apple", "apple", "banana", "", "cherry", ""}
	compactSlice := slices.Compact(slice4)
	fmt.Println("Compact Slice:", compactSlice)
	fmt.Printf("Underlying array after Compact: %q\n", slice4)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	// This should panic because -1 is out of range
	slices.Insert([]int{1, 2, 3}, -1, 4)

	fmt.Println("## NoGoodbye")
}
