package main

import (
	"fmt"
       )

type Point struct{
	X,Y float32
}

func main() {

	fmt.Println("this is first code!")
	
	//基本数据类型
	student := 10 
	fmt.Printf("student :%d \n",student)
	//changeBase(student)
	changeBaseByPoint(&student)
	fmt.Printf("student :%d \n",student)

	//数组
	array := [...]int{1,2,3,4,5}
	for i,v := range array {
		fmt.Printf("array[%d] : %d \n",i,v)
	}
	//changeArray(array)
	changeArrayByPoint(&array)
	for i,v := range array {
		fmt.Printf("array[%d] : %d \n",i,v)
	}


	//切片
	numSlice := []int{1,2,3,4,5,6,7,5}
	for i,v := range numSlice {
		fmt.Printf("numSlice[%d] : %d \n",i,v)
	}
	//changeSlice(numSlice)
	changeSliceByPoint(&numSlice)
	for i,v := range numSlice {
		fmt.Printf("numSlice[%d] : %d \n",i,v)
	}


	//字典
	limap := map[string]string{"a":"A","b":"B","c":"C","d":"D"}
	for k,v := range limap {
		fmt.Printf("limap[%s] : %s \n",k,v)
	}
	changeMap(limap)
	//changeMapByPoint(&limap)
	for k,v := range limap {
		fmt.Printf("limap[%s] : %s \n",k,v)
	}


	//结构体
	/*
	type Point struct{
		X,Y float32
	}
	*/
        var st Point
	fmt.Printf("Point %#v \n",st)
	changeStruct(st)
	//changeStructByPoint(&st)
	fmt.Printf("Point %#v \n",st)
}

// 基本类型通过参数无法改变其值
func changeBase(b int) {
	b = 34
}

// 基本类型通过指针可以改变其值
func  changeBaseByPoint(p *int) {
	*p = 34
}

// 数组通过参数无法改变其值
func changeArray(array [5]int) {
	for i := range array {
		array[i] = 0
	}
}

// 数组通过指针可以改变其值
func  changeArrayByPoint(p *[5]int) {
	for i := range *p {
		p[i] = 78
	}
}

// 切片通过参数改变其值
func changeSlice(s []int) {
	for i := range s {
		s[i] = 78
	}
}

// 切片通过指针改变其值
func changeSliceByPoint(p *[]int){
	for i := range *p {
		(*p)[i] = 33
	}
}

// 字典通过参数修改其值
func changeMap(m map[string]string) {
	for k := range m {
		m[k] = ""
	}
}

// 字典通过指针改变其值
func changeMapByPoint(p *map[string]string) {
	for k := range *p {
		(*p)[k] = ""
	}
}

// 结构体通过参数不能修改其值
func changeStruct(s Point) {
	s.X = 1
	s.Y = 1
}

// 结构体通过指针改变其值
func changeStructByPoint(p *Point) {
	p.X = 1
	p.Y = 1
}

