package main

import (
		"fmt"
		)

func main(){

		//数组初始化
		var  a [3]int
		fmt.Println(a[0])
		fmt.Println(a[len(a) - 1])

		for i, v := range a {
			fmt.Println("i:",i," v:",v)
		}

		for _, v := range a {
			fmt.Println("v :",v)
		}
		
		var q  [3]int = [3]int{1,2,3}
		var r  [3]string = [3]string{"1","2"}
		fmt.Println(q)
		fmt.Println(r)

		//缺省值会进行补充
		m := [3]float32{13,24.5}
		fmt.Println(m)

		//数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型
		//aa := [3]int{1,2}
		//aa = [4]int{4,4}

		type Current  int
		const (
				
				 USD Current = iota
				 EUR
				 GBP
				 RMB
			)
			symbol := [...]string{USD:"$",EUR:"&",GBP:"UF",RMB:"^"}
			fmt.Println(symbol)
			fmt.Println(symbol[RMB])

		bb :=  [...]int{34:-12}
		fmt.Println("bb:",bb)

		cc := [2]int{1,2}
		dd := [2]int{1,2}
		ee := [2]int{3,4}
		fmt.Println(cc == dd,dd == ee,cc == ee)
	
		
 		
}
