package main

import (
	"flag"
	"fmt"
	"log"

	modbus "github.com/advancedclimatesystems/goldfish"
)

func handleRegisters(unitID, start, quantity int) ([]modbus.Value, error) {
	registers := make([]modbus.Value, quantity)
	for i := 0; i < quantity; i++ {
		v, err := modbus.NewValue(i + start)
		if err != nil {
			return nil, modbus.SlaveDeviceFailureError
		}
		registers[i] = v
	}
	return registers, nil
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Modbus Server")

	addr := flag.String("addr", ":502", "address to listen on.")
	flag.Parse()

	server, err := modbus.NewServer(*addr)
	if err != nil {
		log.Fatalf("Failed to start Modbus server: %v", err)
	}

	server.Handle(modbus.ReadHoldingRegisters, modbus.NewReadHandler(handleRegisters))

	server.Listen()
}
