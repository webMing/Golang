package ch5

import (
	"fmt"
)

//MyArray 测试数组内部结构以及特性
func MyArray() {
	const (
		first = iota + 1
		second
		third
		four
		five
	)
	var a = [...]int{first:1,second:2,third:3,four:4,five:5}
	fmt.Println("MyArray:",a)

	aSlice := a[third:four]
	fmt.Println("aSlice:",aSlice)
	
	//修改 aSlice的值,同时也修改 a 的值
	aSlice[0] = 10;
	fmt.Println("MyArray:",a)

	a[third] = 3;
	fmt.Println("MyArray:",a) 

	bSlice := aSlice[0:3]
	fmt.Println("bSlice:",bSlice)
	
}