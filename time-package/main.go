package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	// fmt.Println("Current Time is:", currentTime)
	// formatted := currentTime.Format("dd-mm-yyyy,Day")
	formatted := currentTime.Format("02-01-2006 , Monday")
	fmt.Println(formatted)
}
