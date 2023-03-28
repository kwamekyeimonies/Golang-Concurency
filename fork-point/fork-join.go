package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
		fmt.Println("Done Waiting for the Main exits")
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		Work()
	}()

	// this function is not equal to the code below

	// defer wg.Done()
	// go Work()
	wg.Wait()
}

func Work() {
	fmt.Println("Printing my Work to be done")
}
