package main

import "fmt"

func main() {
	fmt.Println("++++++++++for loop in golang+++++++++++++")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// Infinite loop
	// i := 0
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	i++
	// }

	// range in golang
	numbers := [](int){5, 6, 7, 8, 9}
	for index, value := range numbers {
		fmt.Printf("Index of number is %d, and value of number is %d\n", index, value)
	}
	data := "Pradum Kumar"
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value of data is %c\n", index, char)
	}

}
