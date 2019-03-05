package main

import (
	"fmt"
	"ch5"
)

func main() {
	//ch5.LearnArray()
	//下面操作不会崩溃.
	s := make([]int,3)[3:]
	fmt.Println("len:",len(s),"cap(s):",cap(s))

	fmt.Println()

	ch5.MyArray()
	// ch5.FunValuePassTest()

	//m := [...]int{7:7}[3:] //m :[0,0,0,0,0,0,7]
	//x := m[7:] // x:[0,7]
	//fmt.Println(x[0])
	//fmt.Println(x)
	
}