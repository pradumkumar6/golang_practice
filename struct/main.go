package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	Home   string
	Area   string
	Street string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// First Type
	// var pradum Person
	// pradum.FirstName = "Pradum"
	// pradum.LastName = "Kumar"
	// pradum.Age = 23
	// fmt.Println("Pradum person:", pradum)

	// Second Type
	// person1 := Person{
	// 	FirstName: "Priyanshu",
	// 	LastName:  "Pokhariya",
	// 	Age:       25,
	// }
	// fmt.Println("Person1 is:", person1)

	// Third type by using new keyword
	// var person2 = new(Person)
	// person2.FirstName = "Shreya"
	// person2.LastName = "Pandey"
	// person2.Age = 21
	// fmt.Println("Person2 is:", person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Manish",
		LastName:  "Gupta",
		Age:       24,
	}
	employee1.Person_Contact = Contact{
		Email: "manishgupta60071@gmail.com",
		Phone: "9026506585",
	}
	employee1.Person_Address = Address{
		Home:   "Trilochan",
		Area:   "Jalalpur",
		Street: "Jalalpur",
	}
	fmt.Println("Employee1 is:", employee1)
	fmt.Println("Employee1 email is:", employee1.Person_Contact.Email)
}
