package main

import "fmt"

func main() {
	fmt.Println("++++++++++data conversion++++++++++++")
	var num int = 42
	// fmt.Printf("Type of num is %T\n", num)
	var numConversion string = string(num)
	fmt.Printf("Type of num is %T\n", numConversion)
}
