package structs

import "log"

type myStruct struct {
	FirstName string
	LastName  string
}

func (m *myStruct) GetFullName() string {
	return m.FirstName + " " + m.LastName
}

func ReceiverDemo() {
	var myVar myStruct
	myVar.FirstName = "Munesh"

	myVar2 := myStruct{
		FirstName: "Munesh",
		LastName:  "Kushwah",
	}

	log.Println("First Name:", myVar.FirstName)
	log.Println("First Name:", myVar2.LastName)
	log.Println("Full Name:", myVar2.GetFullName())

}
