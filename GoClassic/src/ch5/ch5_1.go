package ch5

import (
	"fmt"
)

//LearnArray 数组 
func LearnArray() {
	//数组声明
	a1 := [3]int{1,3,4}
	fmt.Println(a1)

	//数组范式1
	a2 := [...]int{3,4,5,6}
	fmt.Println(a2)

	//数组先声明然后再赋值
	a3 := [4]int{}
	fmt.Println(a3)
	a3[3] = 4
	for i := 0;i < len(a3); i++ {
		fmt.Printf("a3[%d]=%d\n",i,a3[i])
	}

	//数组含有下标索引
	const (
		first = iota + 1
		second 
		unkown = 5	
	)
	a4 := [...]int{first:1,second:4,unkown:12}
	for i := 0; i < len(a4); i++ {
		fmt.Printf("a4[%d]= %d \n",i,a4[i])
	}

}

