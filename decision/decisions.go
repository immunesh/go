package decision

import "log"

func Main() {

	var isTrue bool
	myNumber := 100

	isTrue = false

	if isTrue {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}

	cat := "cat2"

	if cat == "cat" {
		log.Println("cat is ", cat)
	} else {
		log.Println("cat is not ", cat)
	}

	if myNumber > 99 && !isTrue {
		log.Println("myNumber is greater than 99 and isTrue is false")
	} else if myNumber < 100 && isTrue {
		log.Println("1")
	} else if myNumber == 101 || isTrue {
		log.Println("2")
	} else if myNumber > 1000 && !isTrue {
		log.Println("3")
	} else {
		log.Println("4")
	}

	myVar := "carscs"

	switch myVar {
	case "car":
		log.Println("myVar is car")
	case "bike":
		log.Println("myVar is bike")
	case "bus":
		log.Println("myVar is bus")
	case "truck":
		log.Println("myVar is truck")
	default:
		log.Println("myVar is not car, bike, bus or truck")
	}
}
