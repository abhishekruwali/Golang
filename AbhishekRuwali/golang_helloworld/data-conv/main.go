package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is", num)
	fmt.Printf("Type of num is %T\n", num)

	// num=num+1.23
	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is ", data)
	fmt.Printf("type of data is %T\n", data)

	// string to integer
	num = 123
	str := strconv.Itoa(num)
	fmt.Printf("integer to string %T\n", str)
	fmt.Println("data is ", str)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string) //return 2 values
	fmt.Printf("integer to string %T\n", number_int)
	fmt.Println("data is ", number_int)

}
