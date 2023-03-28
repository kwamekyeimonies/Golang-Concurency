package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go Worker(wg)
	wg.Wait() 

}

func Worker(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Wait Group without reference &")
}
