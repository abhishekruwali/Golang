package main

import "fmt"

func main() {

	// maps in Go lang

	// studentGrades := make(map[string]int)
	// studentGrades["Prince"] = 34
	// studentGrades["Alice"] = 90
	// studentGrades["Bob"] = 85
	// studentGrades["Charlie"] = 95

	// fmt.Println("Marks of Bob :", studentGrades["Bob"])
	// studentGrades["Bob"] = 95
	// fmt.Println("new marks of Bob", studentGrades["Bob"])

	// // Delete in map
	// delete(studentGrades, "Bob")
	// fmt.Println("Marks of Bob :", studentGrades["Bob"])

	// //  Checking if key is present or not

	// grades, xists := studentGrades["David"]
	// fmt.Println("Grades of David : ", grades)
	// fmt.Println("David exists : ", xists)

	// fmt.Println("Marks of David :", studentGrades["David"])

	// // using for loop

	// for index, value := range studentGrades {
	// 	fmt.Printf("Key is %s and marks is %d\n", index, value)
	// }

	// unordered way store
	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 96,
	}
	for index, value := range person {
		fmt.Printf("Key is %s and marks is %d\n", index, value)

	}

}
