package interfaces

import (
	"fmt"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	Color         string
	numberOfTeeth int
}

func Main() {

	dog := Dog{
		Name:  "Dog",
		Breed: "Labrador",
	}
	gorilla := Gorilla{
		Name:          "Gorilla",
		Color:         "Black",
		numberOfTeeth: 32,
	}

	printInfo(&dog)
	printInfo(gorilla)

}

func printInfo(animal Animal) {
	fmt.Println("Animal Name: ", animal.Says())
	fmt.Println("Animal Number of Legs: ", animal.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberOfLegs() int {
	return 4
}
func (g Gorilla) Says() string {
	return "Ooh Ooh Aah Aah"
}
func (g Gorilla) NumberOfLegs() int {
	return 2
}
func (g Gorilla) NumberOfTeeth() int {
	return 32
}
