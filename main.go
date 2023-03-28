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

	go First_task()
	go Second_Task()
	go Third_Task()
	go Fourth_Task()
	go Fifth_Task()

	time.Sleep(time.Second )
}

func First_task() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("First Task")
}

func Second_Task() {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Second_Task")
}

func Third_Task() {
	// time.Sleep(time.Millisecond * 100)
	fmt.Println("Third_Task")
}

func Fourth_Task() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Fourth_Task")
}

func Fifth_Task() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Fifth_Task")
}
