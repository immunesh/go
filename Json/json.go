package json

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func Main() {

	myJson := `
	[
		{
			"firstName": "Munesh",
			"lastName": "Kushwah",
			"email": "munesh@gmail.com",
			"age": 30
		},	
		{
			"firstName": "John",
			"lastName": "Doe",
			"email": "JohnDoe@gmail.com",
			"age": 25
		}			
	
	]`

	var people []Person
	err := json.Unmarshal([]byte(myJson), &people)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("People: %v\n", people)

	//write json from struct
	var person []Person
	var m1 Person
	m1.FirstName = "Munesh"
	m1.LastName = "Kushwah"
	m1.Email = "munesh6181@gmail.com"
	m1.Age = 30
	person = append(person, m1)
	var m2 Person
	m2.FirstName = "John"
	m2.LastName = "Doe"
	m2.Email = "JohnDoe@gmail.com"
	m2.Age = 25
	person = append(person, m2)

	newJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New Json: %s\n", newJson)
}
