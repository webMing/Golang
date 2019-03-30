package ch8

import (
	"io"
	"time"
	"fmt"
	"net"
	"log"
)

// Serve1 时钟服务1
func Serve1() {
	lisen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := lisen.Accept()
		if err != nil {
			fmt.Println(err)	
			continue
		}
		handle1(con)
	}
}

func handle1(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, time.Now().Format("15-04-05"))
		if err != nil {
			return 
		}	
		time.Sleep(1 * time.Second)
	}
}