package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("We are learning Error Handling")
	res, _ := divide(10, 0)
	fmt.Println("Division of two number is:", res)
}
