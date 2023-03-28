package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
	fmt.Println("Executed Immediately")
}

// Deadlock
// Waiting for something that does not happen
