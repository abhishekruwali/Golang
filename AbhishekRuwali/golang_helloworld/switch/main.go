package main

import "fmt"

func main() {
	// day := 30

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")

	// case 2:
	// 	fmt.Println("Tuesday")

	// case 3:
	// 	fmt.Println("Wednesday")

	// default:
	// 	fmt.Println("Unknown Day")

	// }

	// month := "January"

	// switch month {
	// case "January", "February", "March":
	// 	fmt.Println("Winter")
	// case "April", "May", "June":
	// 	fmt.Println("Spring")
	// default:
	// 	fmt.Println("Other season")
	// }

	// Expression Evaluate
	temperature := 9
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	default:
		fmt.Println("Hot")
	}
}
