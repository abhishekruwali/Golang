package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	// num := 2
	// ptr := &num

	//  fmt.Println("Num has value", num)
	// fmt.Println("Pointer contains", *(ptr))

	// var pointer *int // Default pointer ==nil
	// if pointer == nil {
	// 	fmt.Println("Pointer is not assigned")
	// }

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains :", value)

}
