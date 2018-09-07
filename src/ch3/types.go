package main

import (
		"fmt"
		"unsafe"
		"math"
		)

func main(){
	
	//bool 数据类型, && || ! 操作的类似是bool类型
	fmt.Println("\n \n")
	a := true
	b := false
	fmt.Println("a := true , b:= false")
	fmt.Println("a:",a,"b:",b)
	fmt.Println("a && b:", a && b)
	if !(a && b) {
		fmt.Println("!(a && b) is ", !(a && b))
	}
	fmt.Println("a || b is ", a || b)

    //signed integers
	var c int = 89
	d := 95
	fmt.Println("\n")
	fmt.Println("c is :",c,"d is :",d)

	var j int = 89
	e := 100
	fmt.Println("j is :",j," e is :",e)
	fmt.Println("value of j is",j,"e is",e)
	fmt.Printf("type of j is %T,size of j is %d \n",j,unsafe.Sizeof(j)) //type and size of j  sizeOf() 返回的是字节数
	fmt.Printf("type of e is %T,size of e is %d \n",e,unsafe.Sizeof(e)) //type and size of e  sizeOf() 返回的是字节数

	//floating point types
	k,l := 67,8.97
	fmt.Printf("type of k is %T, type of l is %T \n",k,l)
	fmt.Printf("size of k is %d, size of l is %d \n",unsafe.Sizeof(k),unsafe.Sizeof(l))

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
	//操作在进行移位操作的时候其值不会改变;需要添加左值来接受移位操作后的新值
	var (
		cc uint8 = 0x5
		dd uint8 = 0x3
		)
	fmt.Println("cc dd := 0x5(0b0101),0x3(0b0011)")
	fmt.Printf("cc & dd is :%x \n",cc & dd) //位与 00001 1 1
	fmt.Printf("cc | dd is :%x \n", cc | dd) //位或 0111 7 7
	fmt.Printf("cc ^ dd is :%x \n",cc ^ dd) //异或运算 0110 6 6
	fmt.Printf("cc value is :%#b ; ^cc is : %x \n",cc,^cc) //按位取反 uint8  ^0000 0101 = 1111 1010 =  0xfa
	fmt.Printf("cc value is :%08b ; cc << 1 is :%08b \n",cc,cc << 1) //左移一位 0000 0101 << 1 = 0000 1010
	fmt.Printf("cc value is :%08b ; cc >> 2 is :%08b \n",cc,cc >> 2) //左移一位 0000 0101 >> 2 = 0000 0001
	fmt.Println("\n")
	
	//比较运算
	ee,ff := 2,4
	fmt.Println("ee ff := 2,4")
	fmt.Printf("ee > ff is :%v ; ee < ff is :%v ; ee == ff is : %v ; ee != ff is :%v \n", ee > ff, ee < ff , ee == ff, ee != ff) 
	ee++ //ee++ 是表达式 相当于 ee = ee + 1	
	fmt.Printf("ee++ is :%d \n",ee) 
	fmt.Println("\n")
	
	//类型转换 format
    var jj  int8  = 5
	var kk  int16 = 25
	fmt.Println("jj = ",jj ,"kk =",kk)
	//如果直接进行算术运算因为类型不同而报错
	fmt.Println("jj ++ kk = ",int(jj) + int(kk)) //向数据类型大的int转换保证精度
	fmt.Println("\n")

	//rune 类型
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("type of ascill is : %T, size of ascii is : %d \n",ascii,unsafe.Sizeof(ascii))
	fmt.Printf("%d, %[1]c %[1]q \n",ascii)
	fmt.Printf("%d, %[1]c %[1]q \n",unicode)
	fmt.Printf("%d, %[1]c %[1]q \n",newline)
	fmt.Println("\n")
     
	//22269 是国字的码点, %c rune类型格式化符号;
	fmt.Printf("%c \n",22269)
	fmt.Printf("%c \n",rune(22269))

	//浮点数
	fmt.Printf("max float32 is %f \n",math.MaxFloat32)
	var f float32 = 1 << 24   // 1 << 24 = 1677216
	fmt.Println(f == f + 1,"\n")

	var f1  float64 = 1 << 24
	fmt.Println(f1 == f1 + 1 , "\n")
}
