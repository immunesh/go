package structs

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
	Email       string
	BirthDate   time.Time
}

func Main() {
	user := User{
		FirstName:   "Munesh",
		LastName:    "Kushwah",
		Age:         30,
		PhoneNumber: "1234567890",
		Email:       "munesh6181@gmai.com",
		BirthDate:   time.Date(1993, 1, 1, 0, 0, 0, 0, time.UTC), // Example birth date
	}

	log.Println("User Details:")
	log.Println("First Name:", user.FirstName)
	log.Println("Last Name:", user.LastName)
	log.Println("Age:", user.Age)
	log.Println("Phone Number:", user.PhoneNumber)
	log.Println("Email:", user.Email)
	log.Println("Birth Date:", user.BirthDate)
	log.Println("Birth Date:", user.BirthDate.Format("2006-01-02")) // Format the birth date for display

}
