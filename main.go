package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		// send a value to the channel
		ch <- 42
	}()

	// receive the value from the channel
	value := <-ch

	fmt.Println(value) // Output: 42

	ch <- 48 
	fmt.Println(<-ch)
}

