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

}
