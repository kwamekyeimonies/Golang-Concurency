package main

import (
	"fmt"
)

func main() {

	channel := make(chan string, 4)
	channel <- "Buffer Channel"
	channel <- "Second Channel"
	channel <- "Third Channel"
	channel <- "4th Channel"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
