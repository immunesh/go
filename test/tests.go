package test

import "errors"

func Main() {

	result, err := divide(100, 10)

	if err != nil {
		println(err.Error())
	} else {
		println("Result is ", result)
	}

}

func divide(x, y float32) (float32, error) {

	var result float32

	if y == 0 {
		return 0, errors.New("cannot divide by zero")

	}
	result = x / y
	return result, nil
}
