package main

//Channel Join Point
import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
		fmt.Println("Done Waiting for the Main exits")
	}()

	done := make(chan struct{})

	go func() {
		Work()
		done <- struct{}{}
	}()
	<-done
}

func Work() {
	fmt.Println("Printing my Work to be done")
}
