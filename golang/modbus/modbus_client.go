package main

import (
	"fmt"

	"github.com/goburrow/modbus"
)

func main() {
	// Create a Modbus TCP client
	client := modbus.TCPClient("localhost:502")

	// Read holding registers starting at address 1, count 10
	results, err := client.ReadHoldingRegisters(1, 10)
	if err != nil {
		fmt.Println("Failed to read holding registers:", err)
		return
	}

	// Print the results
	for i, value := range results {
		fmt.Printf("Register %d: %v\n", i+1, value)
	}
}
