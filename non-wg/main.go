package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Number of CPU Cores", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go greeting(i + 1)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Main Funx is done...")
}

func greeting(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello: ", id, "cTask is done")
}
