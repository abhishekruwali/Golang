package main

import "fmt"

func simpleFunction() {
	fmt.Print("Simple Function")
}

// func add(a, b int) int {
// 	return a + b
// }

func add(a int, b int) (result int) {
	result = a + b
	return
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function in go lang")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println("add of two no is:", ans)
	ans1 := multiply(5, 6)
	fmt.Println("mul of two no is", ans1)
}
