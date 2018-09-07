package main

//测试变量逃逸
// go run -gcflags "-m -l" ch6_1.1.go 


import (
	"fmt"
	)


func dummy(b int) int {
	var c int
	c = b
	return c 
}

func void() {

}

func main() {
 	var a int
	void()
	//a escape 因为在要传入Println来打印
	//c escape 因为要在Println打印
	fmt.Println(a,dummy(0))	
}
