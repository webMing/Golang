package main

import (
	"fmt"
	)

type Point struct{X, Y float64}

func main() {
	//类型与类型指针转换,
	p_1 := Point{4,4} 
       	//p_1.ScaledBy_1(0.2) // 类型 -->>  类型 数值不会改变
	p_1.ScaleBy_2(0.5)    // 类型 -->>  指针 数值会改变 
	fmt.Printf("p_1 scaled is %v \n",p_1)	

	p_2 := &Point{4,4}
	p_2.ScaledBy_1(0.5)  // 指针 -->> 类型 数值会改变
	fmt.Printf("p_2 scaled is %v \n",p_2)

	//fmt.Printf("format :%d \n")
	//fmt.Printf("format :%d \n")
	//fmt.Printf("format :%d \n")
	
}

// p是指针还是创建拷贝
func (p Point) ScaledBy_1(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//会不会改变原来的值
func (p *Point) ScaleBy_2(factor float64) {
	p.X *= factor	
	p.Y *= factor
}
