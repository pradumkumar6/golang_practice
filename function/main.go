package main

import "fmt"

func addTwo(a,b int)(int){
	return a+b
}
func main() {
	fmt.Println("We are learning function in Golang.")
	res := addTwo(15,10)
	fmt.Println(res)
}