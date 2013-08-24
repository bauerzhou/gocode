package main

import (
	"fmt"
	"net"
)

var gparam int = 1

func init(){

}

func main(){
	fmt.Println("starting svr")
	listener, err := net.Listen("tcp", "localhost:5300")
	if err != nil {
		fmt.Println("error listening", err.Error())
		return
	}
	
  	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err accepting", err.Error())
			return 
		}
		go doServerStuf(conn)
	}

}

func doServerStuf(conn net.Conn){
	for{
		buf := make([]byte,512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading", err.Error())
			return
		}
                
		fmt.Printf("seq %d received data: %v\n", gparam, string(buf))
		gparam += 1
		conn.Write(buf)
	}
}


