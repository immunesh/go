package variables

import "fmt"

func Main() {

	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")

	var whatToSay string
	whatToSay = "Hello, Go!"
	fmt.Println(whatToSay)
	fmt.Println("This is a simple Go program with a variable.")

	var i int
	i = 42
	fmt.Println("The value of i is:", i)
	fmt.Println("This is a simple Go program with a variable of type int.")
	whatwasSaid, theotherthingthatWasSaid := saySomething()
	fmt.Println("The Function returned", whatwasSaid, "and", theotherthingthatWasSaid)
	fmt.Println("This is a simple Go program with a function that returns two values.")
}

func saySomething() (string, string) {
	// This is a simple function that returns two strings
	return "Something!", "Something else!"
}
