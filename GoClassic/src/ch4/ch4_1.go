package ch4

import (
	"fmt"
)

// FuncNormal 普通的函数
func FuncNormal() {
	fmt.Println("This is first function test")
}
// FuncWithPara 带参数的函数
func FuncWithPara(a,b int) {
	if a  > b {
		fmt.Printf("%d > %d",a,b)
	}else {
		fmt.Printf("%d < %d",a,b)
	}
}
// FuncWithReturnValue 含有返回值的函数
func FuncWithReturnValue(a string,b int)  (r int) {
	fmt.Printf("String:%s,Int:%d \n",a,b)
	return 2
}
// FuncWithMorePara 可变参数
func FuncWithMorePara(a ... int) {
	for _,v := range 	a {
		fmt.Printf("V:%d \n",v)
	}
}

// FuncWithNamedReturnValue 返回多个值
func FuncWithNamedReturnValue() (a,b int ){
	/*
	a = 3
	b = 4
	return	
	*/
	return 3,5
} 

// FuncWithNoName 匿名函数
func FuncWithNoName(f func(int),p int) {
	fmt.Println("Before Func excuted")	
	f(p);
	fmt.Println("Before Func excuted")	
}