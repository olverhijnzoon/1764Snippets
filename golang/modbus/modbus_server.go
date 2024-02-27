package main

import (
	"flag"
	"fmt"
	"log"

	modbus "github.com/advancedclimatesystems/goldfish"
)

// Handle Read Coils and Read Discrete Inputs requests
func handleReadCoils(unitID, start, quantity int) ([]modbus.Value, error) {
	coils := make([]modbus.Value, quantity)
	for i := 0; i < quantity; i++ {
		v, err := modbus.NewValue((i + start) % 2)
		if err != nil {
			return nil, modbus.SlaveDeviceFailureError
		}
		coils[i] = v
	}
	return coils, nil
}

// Handle Read Holding Registers and Read Input Registers requests
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

// Handle Write Single Register and Write Multiple Registers requests
func handleWriteRegisters(unitID, start int, values []modbus.Value) error {
	for i, value := range values {
		fmt.Printf("[%d]: %d\n", i+start, value.Get())
	}
	return nil
}

// Handle Write Single Coil request
func handleWriteCoils(unitID, start int, values []modbus.Value) error {
	if start == 1 {
		return modbus.IllegalAddressError
	}
	return nil
}

func main() {
	addr := flag.String("addr", ":502", "address to listen on.")
	flag.Parse()

	server, err := modbus.NewServer(*addr)
	if err != nil {
		log.Fatalf("Failed to start Modbus server: %v", err)
	}

	server.Handle(modbus.ReadCoils, modbus.NewReadHandler(handleReadCoils))
	server.Handle(modbus.ReadHoldingRegisters, modbus.NewReadHandler(handleRegisters))
	server.Handle(modbus.WriteSingleCoil, modbus.NewWriteHandler(handleWriteCoils, modbus.Signed))
	server.Handle(modbus.WriteSingleRegister, modbus.NewWriteHandler(handleWriteRegisters, modbus.Signed))
	server.Handle(modbus.WriteMultipleRegisters, modbus.NewWriteHandler(handleWriteRegisters, modbus.Signed))

	server.Listen()
}
