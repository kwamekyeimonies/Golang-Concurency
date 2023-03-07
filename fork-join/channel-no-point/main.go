package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Elapsed: ", time.Since(now))
	fmt.Println("Done Waiting")
}

func work() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Some Work needs to be done")
}
