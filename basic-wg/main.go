package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// wg.Add(10)
	// // go func() {
	// // 	fmt.Println("WaitGroup done: 1")
	// // 	defer wg.Done()
	// // }()
	// // go func() {
	// // 	fmt.Println("WaitGroup done: 2")
	// // 	defer wg.Done()
	// // }()
	// // go func() {
	// // 	fmt.Println("WaitGroup done: 3")
	// // 	defer wg.Done()
	// // }()
	// now := time.Now()
	// for i := 0; i < 10; i++ {
	// 	go work(&wg, i+1)
	// }
	// wg.Wait()
	// time.Sleep(100 * time.Millisecond)
	// fmt.Println("Done with Main waitgroup")
	// fmt.Println("Number of Core CPU: ", runtime.NumCPU())
	// fmt.Println("Elapsed: ", time.Since(now))

	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Go routine done")
	}()
	wg.Wait()
	fmt.Println("Execute Immediately after Goroutine")
}

func Work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task: ", id, " is done")
}

// Using wait group
// First create the wait-group
// Add the number of or decide the number of goroutines you want to wait for wg.Add(10)
// Call the wait to wait for the routines and done
