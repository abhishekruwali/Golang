package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World")
	time.Sleep(2000 * time.Millisecond) //simulation some work
	fmt.Println("sayHello function ended successfully")
}

func sayHi() {
	fmt.Println("hi prince")
}

func main() {
	fmt.Println("learning go routine")
	go sayHello() //wait
	sayHi()

	// wait for a moment to allow the goroutine to finish
	time.Sleep(3000 * time.Millisecond)
}
