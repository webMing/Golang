package ch8

import (
	"fmt"
)

var a = new(int)
// Q5 测试分配的物理地址
func Q5() {
	var b = new(int)
	fmt.Printf("New Var A Addr:%p \n", a)
	fmt.Printf("New Var A Addr:%p \n", b)
}