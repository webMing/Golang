package ch8

import (
	"fmt"
	"os"
	"log"
	"io"
	"net"
)

//Diat 网络请求
func Diat() {
	con, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	mustCopy(os.Stdout,con)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_ ,err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
}