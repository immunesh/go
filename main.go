package main

import (
	"fmt"

	//"GO/pointers"
	//"GO/variables"
	//structs "GO/struct"
	//"GO/maps"
	//"GO/decision"
	//"GO/loops"
	//"GO/interface"
	//"GO/channel"
	//json "GO/Json"
	//"GO/test"
	basicweb "GO/basicWebapp"
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

	// Uncomment any of the following sections to run different modules

	////fmt.Println("Running pointers:")
	//pointers.Main()

	//fmt.Println("\nRunning variables:")
	//variables.Main()

	//fmt.Println("\nRunning struct:")
	//structs.Main()

	//fmt.Println("\nRunning struct receiver:")
	//structs.ReceiverDemo()

	//fmt.Println("\nRunning maps and slices:")
	//maps.Main()

	//fmt.Println("\nRunning decision:")
	//decision.Main()

	//fmt.Println("\nRunning loops:")
	//loops.Main()

	//fmt.Println("\nRunning interfaces:")
	//interfaces.Main()

	//fmt.Println("\nRunning channels:")
	//channel.Main()

	//fmt.Println("\nRunning JSON:")
	//json.Main()

	//fmt.Println("\nRunning Tests:")
	//test.Main()

	fmt.Println("\nRunning Basic Web App:")
	basicweb.Main()
}
