package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Number is:", i)
	// }

	// counter := 0
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	counter++
	// 	if counter == 3 {
	// 		break
	// 	}
	// }

	// range
	// numbers := []int{98, 89, 54, 23, 54}
	// for index, value := range numbers {
	// 	fmt.Printf(" index %d value %d \n", index, value)
	// }

	data := "Hello, world!"
	for index, char := range data {
		fmt.Printf("index %d char %c \n", index, char)
	}

}
