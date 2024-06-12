package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Error in Get request", err)
		return
	}
	defer res.Body.Close()
	// fmt.Printf("Type of res is %T\n", res)
	// fmt.Println(res)

	// Read the res body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Reading the file")
		return
	}
	fmt.Println("Response data", string(data))

}
