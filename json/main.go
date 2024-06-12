package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON")
	person := Person{
		Name:    "Pradum",
		Age:     23,
		IsAdult: true,
	}
	fmt.Println("Person is:", person)

	// convert person into JSON  Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in convert to JSON", err)
		return
	}
	fmt.Println("Person data is:", string(jsonData))

	// Decoding(Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling", err)
		return
	}
	fmt.Println("Person data is:", personData)
}
