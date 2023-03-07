package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done Waiting Main")
}

func work() {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Printing some works")
}
