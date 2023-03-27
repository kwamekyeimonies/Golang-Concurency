package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	Food := []string{"Jollof", "Rice and Stew", "Banku and Tilapia", "Fried Rice"}

	for _, my_food := range Food {
		go Attacker(my_food)
	}

	time.Sleep(time.Second * 2)

}

func Attacker(target string) {
	fmt.Println("Throwing Ninja starts", target)
	time.Sleep(time.Second)
}
