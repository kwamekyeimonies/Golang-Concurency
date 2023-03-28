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
	}()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Welcome 1st Person")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Welcome 2nd Person")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Welcome 3rd Person")
	}()

	wg.Wait()

}
