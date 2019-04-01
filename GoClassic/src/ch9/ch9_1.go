/*
 * Go 圣经 第八章课后习题
*/

package ch9

import (
	"os"
	"io"
	"log"
	"net"
	"flag"
)

// Q1 测试
func Q1() {
  client()	
}

var port = flag.String("port", "8000", "Specify the port number, Defalt is 8000")

func client() {
	con, err := net.Dial("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	mustCopy(con)
}

func mustCopy(con net.Conn) {
	if _, err := io.Copy(os.Stdout, con); err != nil {
		log.Println(err)
	}
}