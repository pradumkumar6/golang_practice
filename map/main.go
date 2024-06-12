package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)
	studentGrades["Pradum"] = 90
	studentGrades["Shreya"] = 100
	studentGrades["Manish"] = 95
	studentGrades["Mukesh"] = 92
	// fmt.Println("Marks of Pradum is:", studentGrades["Pradum"])

	// // Upadte the key value in map
	// studentGrades["Pradum"] = 98
	// fmt.Println("New marks of Pradum is:", studentGrades["Pradum"])

	// // Delete any key from map
	// delete(studentGrades, "Manish")
	// fmt.Println("Marks of Manish is:", studentGrades["Manish"])

	// Check if key is exist or not
	// grades, exists := studentGrades["Pradum"]
	// fmt.Println("Marks of Pradum is:", grades)
	// fmt.Println("Pradum exists:", exists)

	grades, exists := studentGrades["Neeraj"]
	fmt.Println("Marks of Pradum is:", grades)
	fmt.Println("Neeraj exists:", exists)
}
