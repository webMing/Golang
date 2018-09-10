package sfunc

import (
		"fmt"
)

//Anony 解释匿名函数
func Anony() {
	anony1()	
}

func anony1() {
	//匿名函数调用 
	func(){
		fmt.Println("I am anomy")
	}() //匿名函数调用

	//匿名函数作为返回值
	c := caculate(1)
	fmt.Printf("c : %d \n",c()) //2
	fmt.Printf("c : %d \n",c()) //3
	fmt.Printf("c address:%p \n",c)

	c1 := caculate(3)
	fmt.Printf("c1 : %d \n",c1()) //4
	fmt.Printf("c1 : %d \n",c1()) //5
	fmt.Printf("c1 address:%p \n",c1)
}

func caculate(v int) func()int {
	//变量捕获
	return func()int {
		v++
		return v 
	}
}