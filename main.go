package main

import (
	"fmt"
	"time"
)

// func main() {
// 	var data int
// 	go func() {
// 		data++
// 	}()
// 	if data == 0 {
// 		fmt.Printf("The value is %v\n", data)
// 	}
// }

// func main(){
// 	var data int
// 	go func ()  {
// 		data++
// 	}()
// 	// time.Sleep(5*time.Second)
// 	if data==0{
// 		fmt.Printf("The value is %v \n",data)
// 	}
// }

// func main() {
// 	var data int
// 	go func() { data++ }()
// 	if data == 0 {
// 		fmt.Println("The value is Zero is 0")
// 	} else {
// 		fmt.Printf("The value is  %v\n", data)
// 	}
// }

func main() {
	now := time.Now()
	done := make(chan struct{})
	go first_task(done)
	go second_task(done)
	go third_task(done)
	go fourth_task(done)
	time.Sleep(time.Second * 1)
	<-done
	<-done
	<-done
	<-done
	fmt.Println("Elapsed: ", time.Since(now))
}

func first_task(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
	done <- struct{}{}
}

func second_task(done chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 2")
	done <- struct{}{}
}

func third_task(done chan struct{}) {
	fmt.Println("Task 3")
	done <- struct{}{}
}

func fourth_task(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 4")
	done <- struct{}{}
}
