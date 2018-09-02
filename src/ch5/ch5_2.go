package main

import (
	"fmt"
	)

func main() {

	//函数座位参数	
	s := []int{1,2,3,4}
	cal(square,s)
	for _,v := range s {
		fmt.Printf("s value :%d \n",v)
	}

	//调用值为nil的函数产生panic
	//var nega func(int)int
	//nega(2) 
	
}

func cal(f func(int)int,s []int) {
	for i := range s {
		s[i] = f(i)
	}
}

func square(n int) int { return n * n }

func negative(n int) int {return -n }

/* 对于错误处理策略
   1. 错误向上传递
   2. 重新尝试失败操作(网络请求)
   3. 仅仅打印错误信息
   4. 忽略错误(忽略错误后对业务逻辑没有影响)
   5. 文件结尾的错误io.EOF
*/

