package main

import "fmt"

func main() {
	x := 4
	// if x > 5 {
	// 	fmt.Println("x is greater than 5.")
	// } else {
	// 	fmt.Println("x is not greater tha 5.")
	// }
	y := 10
	if x < 3 && y > 5 {
		fmt.Println("Hey! Now it is correct.")
	} else {
		fmt.Println("This is not valid😢.")
	}
}
