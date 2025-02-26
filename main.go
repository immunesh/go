package main

import (
	"fmt"
)

func main() {
	// Print a welcome message
	fmt.Println("Hello, World")

	// Ask for the user's name
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	// Greet the user
	fmt.Printf("Hello, %s!\n", name)

	// Ask for two numbers
	fmt.Print("Enter the first number: ")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	var num2 float64
	fmt.Scanln(&num2)

	// Perform a simple addition
	sum := num1 + num2

	// Print the result
	fmt.Printf("The sum of %f and %f is %f\n", num1, num2, sum)
}
