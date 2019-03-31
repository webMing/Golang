package ch8

import (
	"time"
	"io"
	"log"
	"net"
	// "strings"
	// "fmt"
	"flag"
)

var a = new(int)
var tz = flag.String("tz", "Local", "Specified time zone")
var port = flag.String("port", "8000", "Specifiled a port ,Default is 8000")

// Q5 测试分配的物理地址
func Q5() {
	// var b = new(int)
	// fmt.Printf("New Var A Addr:%p \n", a)
	// fmt.Printf("New Var A Addr:%p \n", b)

	// flag.Parse()
	// fmt.Println(*tz)
	// strings.Compare("a", "b")
	server()
}

func server() {
	lisen, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := lisen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handle4(con)
	}
}
func handle4(con net.Conn) {
	defer con.Close()
	for {
		format := "15:03:04 -07"
		location, err := time.LoadLocation(*tz)
		if err != nil {
			log.Println(err)
			break 
		}
		response := time.Now().In(location).Format(format)
		_, er := io.WriteString(con,response)
		if er != nil {
			log.Println(er)
			break 
		}
	}
}
