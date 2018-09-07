package ch2

import (
	"fmt"
)

//交换存储两个整形变量的地址
func swap1(a, b *int) {
	a,b = b,a 
}

func SwapValue1() {
	x,y := 1,3
	swap1(&x,&y)
	fmt.Printf("x:%d,y:%d",x,y)
}