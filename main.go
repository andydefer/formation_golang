package main

import (
	"fmt"
	"math"
)

func main() {
	person := NewPerson("Andy", 83.55, 1.82)

	fmt.Println(person.String())
	fmt.Printf("IMC : %.2f \n", person.IMC())
	fmt.Println("Catégorie :", person.IMCLabel())
}

// =======================
// DOMAIN (STRUCT)
// =======================

type Person struct {
	name   string
	weight float64
	height float64
}

// =======================
// CONSTRUCTOR (idiom Go)
// =======================

func NewPerson(name string, weight, height float64) Person {
	return Person{
		name:   name,
		weight: weight,
		height: height,
	}
}

// =======================
// BUSINESS LOGIC
// =======================

func (p Person) IMC() float64 {
	if p.height <= 0 {
		return 0
	}
	return p.weight / math.Pow(p.height, 2)
}

func (p Person) IMCLabel() string {
	imc := p.IMC()

	switch {
	case imc < 18.5:
		return "Insuffisance pondérale"
	case imc < 25:
		return "Corpulence normale"
	case imc < 30:
		return "Surpoids"
	default:
		return "Obésité"
	}
}

// =======================
// PRESENTATION (clean separation)
// =======================

func (p Person) String() string {
	return fmt.Sprintf(
		"=== Personne ===\nNom: %s\nPoids: %.2f kg\nTaille: %.2f m",
		p.name, p.weight, p.height,
	)
}
