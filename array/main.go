package main

import "fmt"

func main() {
	fmt.Println("We are learning Array Concept in  Golang")
	// First type of array declaration
	var names [5]string
	names[0] = "Pradyum"
	names[1] = "Manish"
	fmt.Printf("Names are %q\n", names)

	// Second Type
	var numbers = [5](int){1, 2, 3, 9, 8}
	fmt.Println("Numbers are:", numbers)
	fmt.Println("Length of array is:", len(numbers))

}
