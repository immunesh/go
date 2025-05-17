package loops

import "log"

func Main() {
	// for loop
	animals := []string{"cat", "dog", "fish", "bird", "lizard"}

	vehicles := make(map[string]string)
	vehicles["car"] = "Toyota"
	vehicles["bike"] = "Honda"
	vehicles["bus"] = "Mercedes"
	for i := 0; i < len(animals); i++ {
		log.Println("Animal is ", animals[i])
	}
	for i, animal := range animals {
		log.Println("Animal is ", animal, " at index ", i)
	}
	for _, animal := range animals {
		log.Println("Animal is ", animal)
	}

	for vehicle, brand := range vehicles {
		log.Println("Vehicle is ", vehicle, " and brand is ", brand)
	}

	var firstline = "Once Upon a midnight dreary"

	for i, l := range firstline {
		log.Println("Character is ", string(l), " at index ", i)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{
		FirstName: "Munesh",
		LastName:  "Kushwah",
		Email:     "munesh6181@gmail.com",
		Age:       30,
	})
	users = append(users, User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "JohnDoe@gmail.com",
		Age:       25,
	})
	users = append(users, User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "JaneDoe@gmail.com",
		Age:       28,
	})
	users = append(users, User{
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "AliceSmith@gmail.com",
		Age:       22,
	})

	for _, user := range users {
		log.Println("User is ", user.FirstName, " ", user.LastName, " with email ", user.Email, " and age ", user.Age)
	}
	for i := 0; i < len(users); i++ {
		log.Println("User is ", users[i].FirstName, " ", users[i].LastName, " with email ", users[i].Email, " and age ", users[i].Age)
	}

}
