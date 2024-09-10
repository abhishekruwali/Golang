package main

import (
	"fmt"
	"time"
)

func main() {
	// "dd-mm-yy"

	//  time is fixed

	// currentTime := time.Now()
	// fmt.Println("Current time ", currentTime)
	// fmt.Printf("Type of currentTime %T\n", currentTime)

	// // fixed time
	// formatted := currentTime.Format("2006-01-02, Monday  15:04:05 PM")
	// fmt.Println("Formatted data:", formatted)

	layout_str := "2006/01/02"
	dateStr := "2023/11/25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time:", formatted_time)

}
