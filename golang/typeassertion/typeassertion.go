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

func (p Particle) String() string {
	return fmt.Sprintf("Particle: %s, Mass: %ekg, Charge: %eC", p.Name, p.Mass, p.Charge)
}

func main() {

	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Type Assertion")

	/*
		In this program, we define an interface Physics which is empty. We also define a struct Particle which represents a particle in physics with properties Name, Mass, and Charge. We then define a method String for the Particle struct which returns a formatted string with the particle's properties.

		In the main function, we create an instance of a Particle, but store it in a variable of type Physics, which is our interface type. We then use a type assertion to check if the Physics object is actually a Particle. If the type assertion succeeds (ok == true), we print the properties of the Particle using the String method we defined. If the type assertion fails (ok == false), we print a message indicating that the Physics object is not a Particle.
	*/

	var physics Physics

	// Elektron mit einer Masse von 9.10938356e-31 kg und einer Ladung von -1.60217662e-19 Coulomb
	physics = &Particle{"Electron", 9.10938356e-31, -1.60217662e-19}

	// Here *Particle is used in a type assertion, where we're checking if the physics interface is actually holding a pointer to a Particle object.
	particle, ok := physics.(*Particle)
	if ok {
		fmt.Println("The properties of the particle are: ")
		fmt.Println(particle.String())
	} else {
		fmt.Println("The physics object is not a Particle")
	}
}
