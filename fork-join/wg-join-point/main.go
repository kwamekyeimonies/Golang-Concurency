package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("Done waiting for main")
	now := time.Now()
	fmt.Println("It took: ", time.Since(now))

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some work to be done")
}
