package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task1(&wg)
	wg.Wait()

	// wg.Add(1)
	// go task2(&wg)
	// wg.Wait()

	wg.Add(1)
	go task3(&wg)
	go task2(&wg)
	wg.Wait()

	wg.Add(2)
	go task3(&wg)
	go task2(&wg)
	wg.Wait()

}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Task 1")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Task 2")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Task 3")
}
