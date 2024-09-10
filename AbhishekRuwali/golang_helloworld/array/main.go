package main

import "fmt"

func main() {
	// fmt.Println("we are learning Array in Golang")

	// var name [5]string
	// name[0] = "prince"
	// name[2] = "Aakash"

	// fmt.Println("Names of person is :", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)
	// Length of string
	// fmt.Println("Length of Numbers are", len(numbers))
	// fmt.Println("value of name at 2 index is", numbers[1])

	// Default values 0 in int case
	// var price [5]int
	// fmt.Println("Price is :", price)

	// //  Default values false in bool case
	// var rose [5]bool
	// fmt.Println("rose is", rose)

	var name [5]string
	name[2] = "prince"
	name[1] = "aakash"
	fmt.Println("Price is :", name)
	// qouted way (string)
	fmt.Printf("Price Array is  %q", name)

	// array dynamic size is not there

}
