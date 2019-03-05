package ch5

import (
	"fmt"
)

// FunValuePassTest 测试函数的值传递
func FunValuePassTest() {
	var a = [...]int{4,5,6,7}
	fmt.Println(a)
	changeValue(a)	
	fmt.Println(a)

	changeVPoint(&a)
	fmt.Println(a)

}

// 把所有的值都改为0; 数组是值传递
func changeValue(a [4]int)(){
	for i := range a {
		a[i] = 0;
	}
}

// 利用指针改变数值;这里传递的是数组的指针
func changeVPoint(adrr *[4]int)() {
	for i:= range adrr {
		adrr[i] = 0
	}
}