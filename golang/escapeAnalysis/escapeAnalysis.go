package main

import (
	"fmt"
)

//go:noinline
func staysOnStack() int {
	// stack0 is created, used, and its value is returned.
	// The variable stack0 itself does not outlive this function.
	// Therefore, it stays on the fast, automatically-cleaned Stack.
	stack0 := 42
	return stack0 * 42
}

//go:noinline
func escapesToHeap() *int {
	// heap1 is created locally, but its memory address is returned.
	// The caller (main) will need access to this memory after escapesToHeap finishes.
	// If y stayed on the stack, it would be destroyed when the function returns.
	// Therefore, the compiler moves it to the Heap.
	heap1 := 42
	return &heap1
}

//go:noinline
func escapesThroughClosure() func() int {
	// heap2 is created locally, but it is captured by the returned anonymous function.
	// Because the closure can be executed long after escapesThroughClosure returns,
	// heap2 must survive beyond this function's scope.
	// Therefore, the compiler moves heap2 to the Heap.
	heap2 := 42
	return func() int {
		heap2++
		return heap2
	}
}

func escapesToGoroutine() chan int {
	// heap3 is created locally, but captured by an asynchronous goroutine.
	// Because the goroutine executes independently and will outlive this function,
	// heap3 cannot be safely cleaned up when escapesToGoroutine returns.
	// Therefore, the compiler moves heap3 to the Heap.
	heap3 := 42
	ch := make(chan int)
	go func() {
		heap3 += 10
		ch <- heap3
	}()
	return ch
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Go Stack vs Heap & Escape Analysis")

	/*
	   In Go, memory is allocated in two main areas:
	   1. The Stack: Very fast, automatic memory management. Used for local variables
	      whose lifetimes are strictly confined to their function.
	   2. The Heap: Slower, managed by Go's Garbage Collector (GC). Used for data
	      that needs to persist across different function calls.

	   Unlike languages like C/C++, you don't explicitly use `malloc` or `free`.
	   Go's compiler performs "Escape Analysis" at compile time. If it detects that
	   a variable's reference "escapes" the function's scope, it moves it to the heap.

	   Use the provided Makefile (`make analyze`) to see exactly what the compiler does!
	*/

	// Example 1: Stack allocation
	// 'val' receives a copy of the integer. No pointers are shared.
	val := staysOnStack()

	// Example 2: Heap allocation
	// 'ptr' receives a pointer to the memory created inside escapesToHeap.
	ptr := escapesToHeap()

	// Example 3: Heap allocation via Closure
	// 'counterFn' receives a function that maintains a reference to 'z'.
	counterFn := escapesThroughClosure()

	// Example 4: Heap allocation via Goroutine
	// 'ch' receives a channel, but the underlying variable 'w' is kept alive by the goroutine.
	ch := escapesToGoroutine()

	// Note: Passing variables to functions like `fmt.Println` often causes them
	// to escape to the heap anyway because they are converted to the `any`
	// (empty interface) type, which requires dynamic allocation.
	fmt.Printf("Stack value: %d\n", val)
	fmt.Printf("Heap pointer value: %d\n", *ptr)
	fmt.Printf("Heap closure value: %d\n", counterFn())
	fmt.Printf("Heap goroutine value: %d\n", <-ch)
}
