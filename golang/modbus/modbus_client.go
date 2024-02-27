package main

import (
	"fmt"

	"github.com/goburrow/modbus"
)

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Modbus Client")

	client := modbus.TCPClient("localhost:502")

	results, err := client.ReadHoldingRegisters(3, 5)
	if err != nil {
		fmt.Println("Failed to read holding registers:", err)
		return
	}

	for i, value := range results {
		fmt.Printf("Register %d: %v\n", i+1, value)
	}
}
