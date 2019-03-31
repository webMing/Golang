package ch7

import (
	"fmt"
	"flag"
)

var para = flag.String("name", "", "display name")
// Q4 测试函数
func Q4() {
	fmt.Println("This is Q4")
	flag.Parse()	
	fmt.Sprintln(*para)
}