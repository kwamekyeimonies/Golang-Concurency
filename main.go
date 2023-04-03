package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Dexoangle"
	go Attacker(evilNinja, smokeSignal)
	smokeSignal <- true
	fmt.Println(<-smokeSignal)

}

func Attacker(target string, mychx chan bool) {
	fmt.Println(target)
	time.Sleep(time.Second)
	mychx <- true
}
