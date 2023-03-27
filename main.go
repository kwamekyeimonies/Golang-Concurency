package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	defer func() {
		fmt.Println("Time taken is: ", time.Since(now))
	}()

	signal_channel := make(chan bool)
	monies := "Daniel Tenkorang"
	go Attacker(monies, signal_channel)
	fmt.Println(<-signal_channel)

	fmt.Println("I Love You")
	signal_channel <- false
	fmt.Println(<-signal_channel)

}

func Attacker(target string, attacked chan bool) {
	fmt.Println("Throwing Ninja starts", target)
	attacked <- true

}
