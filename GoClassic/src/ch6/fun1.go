//Package ch6 中主要学习go中函数相关的内容
package ch6

import (
	"fmt"
)

//F1 简单函数
func F1(){
	fmt.Println("简单函数的展示")
}

//F2 带有一个参数的函数
func F2(name string) {
	fmt.Println("your name is :",name)
}

//F3 一个参数和一个返回值
func F3(name string) (jon string) {
	return name + "add string"	
}

// F4 可变参数
func F4(y ...int) {
	for i,v := range y {
		fmt.Println("index :",i,"value:",v )
	}
}

// F5 匿名函数
func F5(f func(a int)(num int),para int) {
	fmt.Println("Befor excute!")
	fmt.Println(f(para))
	fmt.Println("after excute!")
}