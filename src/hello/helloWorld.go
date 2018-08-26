package main

import "fmt"

func main(){
    
	//hello World!
	fmt.Println("Hello World")
	
	//用var整数变量
	var  a = 1;
	//注意Println 与Printf打印的区别 
	fmt.Println("var a value is :%d",a)
     
	//用var字符串变量
	var  b = "xiaoming"
	fmt.Printf("var b value is :%s \n",b)

	//用var声明的多个字符串的变量
	var c,d = "lixiao","man"
	//Printf格式化不能这么写，Printfln可以这么写.
	fmt.Println("var c value is :",c,"var d vaule is :",d)
	fmt.Printf("var c value is :%s, var d vaule is :%s \n",c,d)
   
	//用var声明的多个不同类型的变量
	var e,f = "linux",34.9
	fmt.Println("var e value is :",e,"var f value is :%d ",f)

	//另一种用var声明过个不同类型的变量的方法
	var(
		j = "A"
		h = 39
		i int
	   )
	fmt.Println(j,h,i)
     
	//用短方式声明的变量形式(short hand declaration)
	k := "short hand declaration type"
	fmt.Printf("var k value is :%s \n",k)

	//用短方式声明的变量中至少有一个是新的
	k,l := "short hand declaration type 1",30
	//Printf格式化不能这么写
	//fmt.Printf("var k value is :%s ",k,"var l value is :%d \n",l)
	fmt.Printf("var k value is :%s , var l value is :%d \n",k,l)
     
}
