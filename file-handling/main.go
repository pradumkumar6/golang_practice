package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Creating a file
	// file, err := os.Create("demo.txt")
	// if err != nil {
	// 	fmt.Println("Error in file creating,", err)
	// 	return
	// }
	// defer file.Close()

	// Write a content in the file
	// content := "Hello from IIT Dhanbad"
	// io.WriteString(file, content)

	// fmt.Println("successfully created")

	// Opening a file
	file, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("Error in file opening", err)
		return
	}
	defer file.Close()

	// Create a buffer to read the file content
	buffer := make([]byte, 1024)

	// Read the file content in buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error in file reading", err)
			return
		}
		// proces the read content
		fmt.Println(string(buffer[:n]))

	}
}
