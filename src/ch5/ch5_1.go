package main

import (
	"fmt"
	"time"
)

const ff  time.Duration = 0

func main() {
	fmt.Printf("%T %v \n",ff,ff)
	fmt.Println("Test func!")
	a := add_type7("3",2)
	fmt.Printf("a value %d \n",len(a))
}

// 函数声明形式
/*

 func function_name(para_list) return_list {
	func_body
 }
*/

func add_type1() {
	fmt.Println("函数没有参数 没有返回值!")
}

func add_type2(a int) {
	fmt.Printf("函数有参数:%d 没有返回值", a)
}

func add_type3(a int, b int) {
	fmt.Printf("函数有参数:%d %d 没有返回值", a, b)
}

func add_type4(a, b int) {
	fmt.Printf("函数有参数:%d %d 没有返回值", a, b)
}

func add_type5(a string, b int) {
	fmt.Printf("函数有参数:%s %d 没有返回值", a, b)
}

func add_type6(a string, b int) string {
	fmt.Printf("函数有参数:%s 有一个返回值 %d", a, b)
	return "ssdfd"
}

func add_type7(a string, b int) (r string) {
	fmt.Printf("函数有参数:%s %d 有一个返回值: %s",a,b,r)
	return 
}
