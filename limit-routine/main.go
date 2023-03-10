package main

import (
	"fmt"
	"sync"
)

type request func()


func main(){

	requests := map[int]request{}
	for i:=0; i<=100;i++{
		f :=func(n int) request{
			return func(){
				fmt.Println("Request",n)
			}
		}
		requests[i]=f(i)
	}
	var wg sync.WaitGroup
	max := 10
	for x:=0;x<len(requests);x+=max{
		for y:=x;y<x+max;y++{
			wg.Add(1)
			go func(r request){
				defer wg.Done()
				r()
			}(requests[y+1])
		}
		wg.Wait()
		fmt.Println(max,"Requests Processed")
	}
}