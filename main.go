package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := make(chan string)
	numRounds := 5
	go Throwing_Start(channel, numRounds)
	// fmt.Println(<-channel)
	for message := range channel {
		fmt.Println(message)
	}
}

func Throwing_Start(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("You scored: %v", score)
	}

}
