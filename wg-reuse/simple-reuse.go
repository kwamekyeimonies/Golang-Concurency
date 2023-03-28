package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Waiting after done")
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

// panic: sync: WaitGroup is reused before previous Wait has returned
