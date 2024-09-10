package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.823

	// fmt.Println("age:", age, "height:", height, "name:", name)
	// fmt.Println("Hello world")

	// fmt.Printf("age is %d", age)
	// fmt.Printf("height is %.3f\n", height)

	// fmt.Printf("Type of name is %T\n", name)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name :%s , Age : %d, Height: %.2f\n", name, age, height)
}
