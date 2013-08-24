package main

import (
 	"fmt"
	"net"
	"runtime"
)

func request(){
	conn, err:= net.Dial("tcp", "localhost:5300")
	if err != nil {
		fmt.Println("conn 5300 fail", err.Error())
		return
	}
	_ , err = conn.Write([]byte("test,...test"))
	read :=true
	buf := make([]byte,512)
	for read {
		_, err = conn.Read(buf)
		fmt.Printf("rsp:%v \n", string(buf))
	}
 //      conn.Close()
}

func main(){
	runtime.GOMAXPROCS(2)
     	for i:=0; i < 100; i++ {
		go request()
	}


}
