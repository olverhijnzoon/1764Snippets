package main

import (
	"fmt"
)

type Physics interface{}

type Particle struct {
	Name   string
	Mass   float64
	Charge float64
}

type Wave struct {
	Frequency float64
	Amplitude float64
}

func (p Particle) String() string {
	return fmt.Sprintf("Particle: %s, Mass: %ekg, Charge: %eC", p.Name, p.Mass, p.Charge)
}

func (w Wave) String() string {
	return fmt.Sprintf("Wave: Frequency: %eHz, Amplitude: %em", w.Frequency, w.Amplitude)
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Type Assertion")

	/*
		In this program, we define an interface Physics which is empty. We also define a struct Particle which represents a particle in physics with properties Name, Mass, and Charge. We then define a method String for the Particle struct which returns a formatted string with the particle's properties.

		In the main function, we create an instance of a Particle, but store it in a variable of type Physics, which is our interface type. We then use a type assertion to check if the Physics object is actually a Particle. If the type assertion succeeds (ok == true), we print the properties of the Particle using the String method we defined. If the type assertion fails (ok == false), we print a message indicating that the Physics object is not a Particle.

		This is a modified version of a 42Snippets snippet.

		Check: https://go.dev/blog/laws-of-reflection
	*/

	// Demonstrating a successful type assertion
	var physics Physics = &Particle{"Electron", 9.10938356e-31, -1.60217662e-19}
	if particle, ok := physics.(*Particle); ok {
		fmt.Println("Successful type assertion:", particle)
	} else {
		fmt.Println("Failed type assertion for Particle")
	}

	// Introducing another type that implements Physics
	physics = &Wave{Frequency: 1e14, Amplitude: 1e-2}

	// Attempting a failed type assertion
	if particle, ok := physics.(*Particle); ok {
		fmt.Println("Successful type assertion:", particle)
	} else {
		fmt.Println("Failed type assertion: Physics object is not a Particle")
	}

	// Using type switch to handle different possible types stored in the interface
	switch v := physics.(type) {
	case *Particle:
		fmt.Println("Found a Particle:", v)
	case *Wave:
		fmt.Println("Found a Wave:", v)
	default:
		fmt.Println("Unknown type")
	}
}
