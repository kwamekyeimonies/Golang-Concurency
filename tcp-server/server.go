package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)


func main(){
	li,err := net.Listen("tcp",":8090")
	if err != nil{
		log.Fatalf("Could not create listener: %v\n",err)
	}

	var connections int32

	for {
		conn,err := li.Accept()
		if err !=nil{
			continue
		}
		connections++

		go func(con net.Conn){
			defer func()  {
				_ = con.Close()
				atomic.AddInt32(&connections,-1)	
			}()

			if atomic.LoadInt32((&connections)) >3{
				return
			}

			time.Sleep(time.Second)
			_,err:=conn.Write([]byte("Success"))
			if err != nil{
				log.Fatalf("Could Not Create Listener: %v\n",err)
			}
		}(conn)

	}
	
}