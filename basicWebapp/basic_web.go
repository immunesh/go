package basicweb

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page!")
	fmt.Println("Endpoint Hit: HomePage")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	sum := addValues(5, 10)
	fmt.Fprintf(w, fmt.Sprintf("Sum is %d", sum))
	fmt.Println("Endpoint Hit: AboutPage")
}

func addValues(x, y int) int {
	return x + y
}

func DividePage(w http.ResponseWriter, r *http.Request) {
	result, err := divide(100, 0)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, fmt.Sprintf("Result is %f", result))
	}
	fmt.Println("Endpoint Hit: DividePage")
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result = x / y
	return result, nil
}

func Main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	// })
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/divide", DividePage)
	_ = http.ListenAndServe(":8080", nil)
}
