package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	totoal, max := 10, 3
	var wg sync.WaitGroup

	for i := 0; i < totoal; i = i + max {
		limit := max
		if i+max > totoal {
			limit = totoal - i
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8090")
				if err != nil {
					log.Fatalf("Could not dial: %v", err)
				}

				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("Could not dial: %v", err)
				}

				if string(bs) != "Success" {
					log.Fatalf("Request Error, Request: %d", i+1+j)
				}

				fmt.Printf("Request: %d: Success \n ", i+1+j)

			}(j)
		}

		wg.Wait()

	}
}
