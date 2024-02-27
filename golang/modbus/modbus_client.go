package main

import (
	"fmt"

	"github.com/goburrow/modbus"
)

func main() {
	// Create a Modbus TCP client
	client := modbus.TCPClient("localhost:502")

	results, err := client.ReadHoldingRegisters(0, 20)
	if err != nil {
		fmt.Println("Failed to read holding registers:", err)
		return
	}

	// Print the results
	for i, value := range results {
		fmt.Printf("Register %d: %v; ", i+1, value)
	}
}
