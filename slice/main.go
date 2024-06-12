package main

import "fmt"

func main() {
	numbers := []int{1, 2, 5, 8, 7}
	numbers = append(numbers, 10, 12, 27, 16, 19, 20)
	fmt.Println("Numbers are:", numbers)
	fmt.Println("Length is:", len(numbers))
}
