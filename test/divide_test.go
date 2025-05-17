package test

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"divide-by-zero", 100.0, 0.0, 0.0, true},
}

func TestMain(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := divide(test.divident, test.divisor)

			if test.isErr && err == nil {
				t.Errorf("Expected an error, got result: %f", result)
			} else if !test.isErr && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			} else if result != test.expected {
				t.Errorf("Expected %f, got %f", test.expected, result)
			}
		})
	}
}

// func TestDivide(t *testing.T) {
// 	result, err := divide(100, 10)

// 	if err != nil {
// 		println(err.Error())
// 	} else {
// 		println("Result is ", result)
// 	}
// }
// func TestDivideByZero(t *testing.T) {
// 	result, err := divide(100, 0.0)

// 	if err == nil {
// 		t.Errorf("Expected an error, got result: %f", result)
// 	} else {
// 		println("Error is ", err.Error())
// 	}
// }
