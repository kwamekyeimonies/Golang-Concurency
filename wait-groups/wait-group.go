package main

//Wait-Group
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
		fmt.Printf("Number of Cores %v\n", runtime.NumCPU())
		fmt.Println("Without Waitgroup")
	}()

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Work(i+1, &wg)
	}

	wg.Wait()
}

func Work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(100 * time.Millisecond)
	fmt.Println("Number: ", id)
}
