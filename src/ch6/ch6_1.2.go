package main

// 测试变量逃逸
// go run -gcflags "-m -l" ch6_1.2.go

import (
	"fmt"
	)

type Data struct{}

func dummy() *Data{
	var c Data
	return &c
}


func main() {
	fmt.Println(dummy())
}
