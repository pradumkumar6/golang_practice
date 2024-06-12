package main

import "fmt"

func main() {
	var num int
	num = 5
	var ptr *int
	ptr = &num
	fmt.Println("Value of num is:", num)
	fmt.Println("Address of num is:", ptr)
}
