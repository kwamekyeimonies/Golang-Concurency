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

	done := make(chan struct{})

	go First_task(done)
	go Second_Task(done)
	go Third_Task(done)
	go Fourth_Task(done)
	go Fifth_Task(done)

	<-done
	<-done

}

func First_task(done chan struct{}) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("First Task")
	done <- struct{}{}
}

func Second_Task(done chan struct{}) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("Second_Task")
	done <- struct{}{}
}

func Third_Task(done chan struct{}) {
	// time.Sleep(time.Millisecond * 100)
	fmt.Println("Third_Task")
	done <- struct{}{}
}

func Fourth_Task(done chan struct{}) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Fourth_Task")
	done <- struct{}{}
}

func Fifth_Task(done chan struct{}) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Fifth_Task")
	done <- struct{}{}
}
