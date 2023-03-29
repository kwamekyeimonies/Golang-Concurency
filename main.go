package main


import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
		fmt.Println("Without Waitgroup")
	}()

	for i := 0; i < 10; i++ {
		go Work(i + 1)
	}

	time.Sleep(100 * time.Millisecond)
}

func Work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Number: ", id)
}
