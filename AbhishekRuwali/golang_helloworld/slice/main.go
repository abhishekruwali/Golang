package main

import "fmt"

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10, 12, 13, 14, 15, 16, 17, 18, 19)
	// fmt.Println("Number :", numbers)
	// fmt.Printf("Number has data type : %T\n", numbers)
	// fmt.Println("Length :", len(numbers))

	// 0 length slice
	// name := []string{}
	// fmt.Println("name :", name)

	// k := []int{}
	// fmt.Println(k)

	// numbers := []int{1, 2, 3}

	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity", cap(numbers))

	// using make

	// numbers := make([]int, 3, 5)

	// numbers = append(numbers, 4)
	// numbers = append(numbers, 98)
	// numbers = append(numbers, 6)

	// fmt.Println("Slice:", numbers)
	// fmt.Println("Length", len(numbers))
	// fmt.Println("capacity", cap(numbers))

	// var stock=[]string  Error

	var numbers = make([]string, 0)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("capacity", cap(numbers))

}
