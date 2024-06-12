package main

import "fmt"

func main() {
	fmt.Println("+++++++++++++Switch statement++++++++++++")
	day := 8
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunady")
	default:
		fmt.Println(("Not a day"))

	}

}
