package main

import "fmt"

func main() {

	x, y := make(chan string), make(chan string)
	go Elect(x, "Daniel")
	go Elect(y, "Tenkorang")

	select {
	case message := <-x:
		fmt.Println(message)
	case message := <-y:
		fmt.Println(message)
	}

}

func Elect(ninja chan string, message string) {
	ninja <- message
}
