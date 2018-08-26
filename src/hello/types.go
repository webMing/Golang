package main

import (
		"fmt"
		"unsafe"
		)

func main(){
	
	/*bool*/
	a := true
	b := false
	fmt.Println("a:",a,"b:",b)

	c := a && b
	fmt.Println("c:",c)
    
	//逻辑非
	if !c {
		fmt.Println("c is true")
	}
	d := a || b
	fmt.Println("d:",d)

    /*signed integers*/
	var h int = 89
	f := 95
	fmt.Println("Value of h is ",h,"and f is",f)

	var j int = 89
	e := 100
	fmt.Println("value of j is",j,"and e is",e)
	fmt.Printf("type of j is %T,size of j is %d \n",j,unsafe.Sizeof(j)) //type and size of j  sizeOf() 返回的是字节数
	fmt.Printf("type of e is %T,size of e is %d \n",e,unsafe.Sizeof(e)) //type and size of e  sizeOf() 返回的是字节数

	//floating point types
	k,l := 5.67,8.97
	fmt.Printf("type of k is %T, type of l is %T",a,b)
	fmt.Printf("size of k is %d, type of l is %d",unsafe.Sizeof(a),unsafe.Sizeof(b))

	sum := k + l
	diff := k - l
	fmt.Println("sum",sum,"diff",diff)

	no1,no2 := 56,80
	fmt.Println("sum",no1 + no2,"diff",no1 - no2)

	// 不同数据类型之间进行运算(不运行)
	//fmt.Printfln("mut cacu", k - no1)
	
	//逻辑运算
	aa,bb := true,false
	fmt.Println("\n")
	fmt.Println("aa,bb := true,false")
	fmt.Println("aa && bb is :",aa && bb)
	fmt.Println("aa || bb is :", aa || bb)
	fmt.Println("!aa:",!aa)
	fmt.Println("!bb:",!bb)
	fmt.Println("\n")
     
	/*  && || ！用于操作bool 类型的数值,不能用于整型
	cc,dd := 10,0
	fmt.Println("\n")
	fmt.Println("cc,dd := 10,0")
	fmt.Println("cc && dd is :",cc && dd)
	fmt.Println("cc || dd is :", cc || dd)
	fmt.Println("!cc:",!cc)
	fmt.Println("!dd:",!dd)
	fmt.Println("\n")
	*/
	
	//位运算 一般进行位运算的数是无符号数;如果对有符号数进行位运算要注意
	//cc,dd := 0x5,0x3
	//%b 是一种打印二进制的格式化字符串
	var (
		cc uint8 = 0x5
		dd uint8 = 0x3
		)
	fmt.Println("\n")
	fmt.Println("cc dd := 0x5(0b0101),0x3(0b0011)")
	fmt.Printf("cc & dd is :%x \n",cc & dd) //位与 00001 1 1
	fmt.Printf("cc | dd is :%x \n", cc | dd) //位或 0111 7 7
	fmt.Printf("cc ^ dd is :%x \n",cc ^ dd) //异或运算 0110 6 6
	fmt.Printf("cc value is :%#b ; ^cc is : %x \n",cc,^cc) //按位取反 uint8  ^0000 0101 = 1111 1010 =  0xfa
	fmt.Printf("cc value is :%#b ; cc << 1 is :%#b \n",cc,cc << 1) //左移一位 0000 0101 << 1 = 0000 1010
	fmt.Printf("cc value is :%#b ; cc >> 2 is :%#b \n",cc,cc >> 2) //左移一位 0010
	fmt.Println("\n")
	
	
	//比较运算
	ee,ff := 2,4
	fmt.Println("\n")
	fmt.Println("ee ff := 2,4")
	fmt.Printf("ee > ff is :%v ; ee < ff is :%v ; ee == ff is : %v ; ee != ff is :%v \n", ee > ff, ee < ff , ee == ff, ee != ff) 
	ee++ //ee++ 是表达式 相当于 ee = ee + 1	
	fmt.Printf("ee++ is :%d \n",ee) 
	fmt.Println("\n")
	
}
