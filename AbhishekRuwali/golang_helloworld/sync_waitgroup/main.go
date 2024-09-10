package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("worker %d started \n", i)
	// some rask is happening
	fmt.Printf("worker %d end \n", i)

}

// worker(1)
// worker(2)
// worker(3)

func main() {
	// fmt.Println("Explore goroutine started")

	// Start 3 worker goroutine
	// wg variable checklist track which goroutines are added in our checklist
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) //Increment the Wait group counter
		go worker(i, &wg)
	}
	// Wait for all workers to finsh
	wg.Wait()
	fmt.Println("worker task completed")
}
