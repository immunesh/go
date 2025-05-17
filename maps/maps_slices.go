package maps

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func Main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Munesh",
		LastName:  "Kushwah",
	}
	myMap["me"] = me

	// Slices
	var mySlice []int
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 4)

	sort.Ints(mySlice) // Sort the slice in ascending order

	log.Println("Slice:", mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println("Numbers:", numbers[0:5]) // Slice from index 0 to 4 (5 is exclusive)

}
