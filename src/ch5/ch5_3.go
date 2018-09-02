package main

import (
	"fmt"
	)

func main() {

	fmt.Println("----------------")

 	//函数嵌套调用
	strF := func(s string)string{return handleStr(s)}
	fmt.Printf("Hi :%s",strF("lixiao"))	

	//普通函数的多次调用
         a := addTwice(1)     // x = 2 p:1
	 fmt.Printf("a value : %d \n",a)

	 a = addTwice(2)      // x = 2 p:2			
	 fmt.Printf("a value : %d \n",a)

	 a = addTwice(0)      // x = 2, p:0		
	 fmt.Printf("a value : %d \n",a)
	
        //函数作为函数值
	 bf := addTwice
         b := bf(1)         // x = 2 p:1
	 fmt.Printf("b value : %d \n",b)

	 b = bf(2)	    // x = 2 p:2		
	 fmt.Printf("b value : %d \n",b)

	 b = bf(0)	   //  x = 2 p:0 	
	 fmt.Printf("b value : %d \n",b)
	

	//函数值的引用不带参数
	f := squares()
	fmt.Printf("f : %d \n", f()) //1
	fmt.Printf("f : %d \n", f()) //4
	fmt.Printf("f : %d \n", f()) //9
	fmt.Printf("f : %d \n", f()) //16
	
	//函数值的引用带参数
	q := squares1(2)
	fmt.Printf("q : %d \n", q()) //3
	fmt.Printf("q : %d \n", q()) //4
	fmt.Printf("q : %d \n", q()) //5

	//简单测试变量捕获	
	test := 1
	testf := func()int{ test++; return test}
	fmt.Printf("testf :%d testValue:%d \n",testf(),test) //2 是引用不是值的拷贝
	test = 30
	fmt.Printf("testf :%d  testValue:%d \n",testf(),test) //31 是引用不是值的拷贝
	
	//变量捕获测试
	 //captureVar := 1
	 //captureF := func	
	
	//陷阱测试
	nums := [...]int{1,2,3,4,5,6,7,8}
	var nums_s []func()
	for index,value := range nums {
		value := value
		nums_s = append(nums_s,func(){
			fmt.Printf("index: %d value: %d \n",index,value)
		})	
	}
	for _, f := range nums_s {
		f() //函数调用
	}
	
	//陷阱测试2
	fmt.Println("\n")
	var score []func()
	for index,value := range nums {
		value := value
		score = append(score,func(){
			fmt.Printf("index: %d value:%d \n",index,value)
		})
	}
	for _, f := range score {
		f()
	}
	fmt.Println("----------------")
}

func handleStr(s string) string {
	return fmt.Sprintf("I am %s \n",s)
}

//捕获变量没有参数
func squares() func()int {
	var x int
	return func()int{
		x++
		return x * x
	} //捕获变量x
}

//捕获变量含有参数
func squares1(n int) func()int {
	var x int
	n++
	return func()int{
		x++
		return x + n
	}
}

func  addTwice(n int) int {
	var x = 1
	x++
	fmt.Printf("x value: %d \n",x)
	return n + x
}

//func printfFunc(f func())
