package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split the data
	// fruits := "Apple,Banana,Grapes,Papaya"
	// data := strings.Split(fruits, ",")
	// fmt.Println(data)

	// Count the data
	// numbers := "One,Two,Three,Four,One,Two,Two,Two,Two"
	// count := strings.Count(numbers, "Two")
	// fmt.Println("Count of two is:", count)

	// Trim the spaces
	data := "      Hello! IIT Dhanbad        "
	trimData := strings.TrimSpace(data)
	fmt.Println(trimData)
}
