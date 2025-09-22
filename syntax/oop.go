package syntax

import (
	"fmt"
)

// Base struct (like a "class")
type Animal struct {
	Name string
}

// Method attached to Animal
func (a Animal) Speak() {
	fmt.Printf("%s makes a sound.\n", a.Name)
}

// Dog struct "inherits" (embeds) Animal
type Dog struct {
	Animal
	Breed string
}

// Overriding Speak() method for Dog
func (d Dog) Speak() {
	fmt.Printf("%s the %s barks!\n", d.Name, d.Breed)
}

// Interface for polymorphism
type Speaker interface {
	Speak()
}

func main() {
	// Create Animal object
	animal := Animal{Name: "Generic Animal"}
	animal.Speak()

	// Create Dog object
	dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
	dog.Speak()

	// Demonstrating polymorphism
	var s Speaker

	s = animal
	s.Speak() // calls Animal's Speak

	s = dog
	s.Speak() // calls Dog's Speak
}
