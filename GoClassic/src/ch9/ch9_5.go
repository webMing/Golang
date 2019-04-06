package ch9

import (
	"fmt"
)

// Q5 测试匿名函数的返回值
func Q5(){
    f := func()(n int){
		return 3
	}
	 c := f()
	  fmt.Println(c)
//    fmt.Println(f())
	 fmt.Println("Hello World")	
}