package main

import (
		"fmt"
		"math"
		"unsafe"
		"math/cmplx"
	   )

func main() {
	
	fmt.Println("\n")
	
	//数值之间类型转换
	fmt.Println("")
	var apples int = 32
	var oranges int8 = 4
	fmt.Println("apples + oranges = ", apples + int(oranges))
	fmt.Println("\n")
	
	var aa  = 3.14
	var bb  = int(aa) 
	fmt.Println("aa is ",aa,"bb is ",bb)
	var cc = 1.99
	fmt.Println("cc is ",int(cc))
	fmt.Println("\n")
     
	/*
    var dd = 1e1000
	var ee = int(dd)
	fmt.Println("ee is",ee) //因为dd 是一个很大的数,可能超出ee int 的范围，这样做很危险
	fmt.Println("\n")
	*/
 
	//浮点数
	for ff := 0; ff < 8 ;ff++ {
		fmt.Printf("ff = %d  e(ff) = %8.3f\n",ff,math.Exp(float64(ff))) //为什么使用float64() ? Exp 的参数要求是float64
	}

	//特殊值
	fmt.Println("")
	var  jj float64
	fmt.Println("jj = ",jj)
	fmt.Println("jj, -jj , jj/jj , 1/jj ,-1/jj",jj,-jj,jj/jj,1/jj,-1/jj)
	fmt.Println("")
	/*
	var ff  = 3.14
	var gg  = int(aa) 
	fmt.Println("aa is ",aa,"bb is ",bb)
	var cc = 1.99
	fmt.Println("cc is ",int(cc))
	fmt.Println("\n");
	*/

	//复数
	var  kk  = complex(1,2)
	fmt.Printf("type of kk is : %T, size of kk is :%d \n",kk,unsafe.Sizeof(kk))
 
	var  ll =  complex(3,4)
	fmt.Println("kk * ll = ",kk * ll)
	fmt.Println("real :",real( kk * ll))
	fmt.Println("imag :",imag(kk * ll))
	fmt.Println("cmplx.Sqrt(-1) :",cmplx.Sqrt(-1))
	fmt.Println("\n")

	//字符串
	var  mm = "你好我的名字叫小王"
	var  subString = mm[3:]
	fmt.Println("subString:",subString)

	var defaultName = "Sam"
	type myString  string
	var customName myString = "Sam"
	customName = myString(defaultName)
	fmt.Println(customName + "\n")
}
