package pointers

import "fmt"

func Main() {
	var myString string
	myString = "Go Green"
	fmt.Println("myString is set to ", myString)
	changeUsingValue(myString)
	fmt.Println("Right After the func call myString is set to ", myString)
	changeUsingPointer(&myString)
	fmt.Println("Right After the func call myString is set to ", myString)

}

func changeUsingPointer(str *string) {
	newValue := "Go Blue"
	*str = newValue
	fmt.Println("Inside changeUsingPointer, myString is set to ", *str)
}

func changeUsingValue(str string) {
	newValue := "Go Red"
	str = newValue
	fmt.Println("Inside changeUsingValue, myString is set to ", str)
}
