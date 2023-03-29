package main

import (
	"fmt"
	"sync"
)

type Request func()

func main() {

	requests := map[int]Request{}
	for i := 0; i <= 100; i++ {
		f := func(n int) Request {
			return func() {
				fmt.Println("Request", n)
			}
		}
		requests[i] = f(i)
	}

	var wg sync.WaitGroup

	max := 10
	for i := 0; i < len(requests); i += max {
		for j := i; j < i+max; j++ {
			wg.Add(1)

			go func(r Request) {
				defer wg.Done()
				r()
			}(requests[j])
		}
	}
	wg.Wait()

}
