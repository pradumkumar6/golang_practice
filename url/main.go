package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	fmt.Printf("Type of url is %T\n", myUrl)
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error in parse the url", err)
		return
	}
	fmt.Printf("Type of url is %T\n", parsedUrl)

}
