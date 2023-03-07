package main

import (
	"log"
	"net"
	"sync/atomic"
)

func main(){
	li,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalf("Could not Create Listener %v\n",err)
	}

	var connections int32

	for {
		conn,err := li.Accept()
		if err != nil{
			continue
		}
		connections++
		go func(){
			defer func(){
				_=conn.Close()
				atomic.AddInt32((&connections,-1))
			}()
			if atomic.
		}()
	}
}