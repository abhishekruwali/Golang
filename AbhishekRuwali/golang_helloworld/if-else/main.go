package main

import "fmt"

func main() {
	// fmt.Printf("how ARE FDBF")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is smaller than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("z is smaller than 5")
	}

}
