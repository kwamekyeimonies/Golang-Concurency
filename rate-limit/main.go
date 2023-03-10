package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)


func main(){
	total,max := 10,3
	var wg sync.WaitGroup
	for i:= 0; i< total; i+= max{
		limit := max
		if i+max > total{
			limit = total - 1
		}

		wg.Add(limit)
		for j:=0; j<limit; j++{
			go func (j int)  {
				defer wg.Done()
				conn,err := net.Dial("tcp",":8080")
				if err != nil{
					log.Fatalf("Could not dial: %v",err)
				}
				bs,err := ioutil.ReadAll(conn)
				if err != nil{
					log.Fatalf("Could not Read from connection %v",err)
				}
				if string(bs) != "success"{
					log.Fatal("request error, request : %d",i+1+j)
				}
				fmt.Printf("Request %d: Success\n",i+1+j)
			}(j)
		}
		wg.Wait()
	}
	
}