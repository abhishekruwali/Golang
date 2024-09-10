package main

import "fmt"

// apne aap ma hi data type ha
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
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Prince",
		LastName:  "kumar",
		Age:       18,
	}

	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "0987654321"

	employee1.Person_Address = Address{
		House: 101,
		Area:  "Ranchi",
		State: "Jharkhand",
	}

	fmt.Println("Employee 1 :", employee1)

	// 1 method
	// var prince Person
	// fmt.Println("Prince Person :", prince) // " " " " 0
	// prince.FirstName = "Prince"
	// prince.LastName = "Agarwal"
	// prince.Age = 24
	// fmt.Println("Prince Person :", prince)

	// // 2 method
	// person1 := Person{
	// 	FirstName: "Aakash",
	// 	LastName:  "Singh",
	// 	Age:       26,
	// }
	// fmt.Println("Person 1 :", person1)

	// // new keyword
	// var person2 = new(Person) //pointer
	// person2.FirstName = "Abhishek"
	// person2.LastName = "Ruwali"
	// person2.Age = 26

	// fmt.Println("Perosn 2:", person2.FirstName)

}
