package ch8

import (
	"io"
	"fmt"
	"net"
	"log"
	"time"
)

// Server 服务器端监听
func Server() {
	lisen, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for {
		acc, err := lisen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handle(acc)
	}

}

func handle(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().Format("15-04-05\n"))
		if err != nil {
			continue	
		}
		time.Sleep(1 * time.Second)
	}
}