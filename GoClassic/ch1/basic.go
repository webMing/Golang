package main

import "fmt"

//变量定义

func  variableZeroValue(){
	/*
	1 变量默认有初始值
	2 对于空字符串可以通过printf中的%q(quotation 引号)格式化输出
	*/
	var a int;
	var s string;
	fmt.Printf("%d %q",a,s)
}

func variableInitialValue() {
	// 变量书写方式之一
	var a int = 3
	var s string = "abc"
	fmt.Println(a,s);	
	//  变量类型会自动推导
	var c = 3
	var d = "abc"
	fmt.Println(c,d);
	
	// 同时定义多个变量
	var e,f int = 2,4
	fmt.Println(e,f)
	
	// 同时定义多个不同类型的变量
	var g,h = 3,"yy"
	fmt.Println(g,h);
	
	// 同时使用var定义更多的变量
	var (
		x = 4
		y = 5
		z = "z"
	)
	fmt.Println(x,y,z)
}

func variableShorter() {
	a,b,c := 2,"b",false
	fmt.Println(a,b,c)	 
}

//在函数外面定义变量;不能使用 shorter mode；只能使用 var
var aa = 33
var ss = "kkk"
var bb = true 

var (
	cc = 4
	mm = 55
	dd = true
)	


func main(){
	// fmt.Println("Hello World!")
	// variableZeroValue();
	// variableInitialValue()
	variableShorter()
}