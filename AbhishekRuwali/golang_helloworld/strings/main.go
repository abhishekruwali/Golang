package main

import (
	"fmt"
	"strings"
)

func main() {

	// Splitting strings
	data := "apple.orange.banana.Prince"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	// count the no of occurence
	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count:", count)

	// Trim function trim in starting and ending
	str = "        Hello, go!     "
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimeed", trimmed)

	// Join function
	str1 := "Prince"
	str2 := "Agarwal"
	result := strings.Join([]string{str1, str2}, " aur bhai ")
	fmt.Println("result:", result)

}
