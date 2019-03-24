package ch8

import (
		"math"
		"fmt"
)

//Point 可导出的自定义类型
type Point struct {X,Y float64}

//Distance Point类型的可导出方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

//Distance 包的可以导出的方法,测试两个点之间的距离
func Distance(p,q Point) float64 {
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

//A1T 测试
func A1T() {
	p := Point{1,2}
	q := Point{4,6}
	///两个函数之间的距离
	fmt.Printf("采用自定义类型的方法计算距离:%.1f \n", p.Distance(q) )
	/// 计算两个点之间的距离采用包的方法
	fmt.Printf("采用包的方法计算距离:%.1f \n",Distance(p,q))
}

//Path 点集之和
type Path []Point 
//Distance Path各个点之间的距离
func (p Path) Distance() float64 {
	var sum float64 
	for i,value := range  p {
		if i > 0 {
			sum += value.Distance(p[i - 1]) 
		}	
	}
	return sum
}

//A2T 测试2
func A2T() {
	path := Path{
		{1,1},
		{5,4},
		{5,1},
		{1,1},
	}
	sum := path.Distance()
	fmt.Printf("sum Distance bettwn:%.2f", sum)
}


//Scale 对Point进行缩放
func (p *Point) Scale(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor

}

//Scale2 对Point进行Scale,这里的操作不会产生任何影响
func (p Point)Scale2(factor float64) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

//A3T 测试3
func A3T() {
	p := Point{1,2}
	fmt.Printf("befor factor: %v \n",p)
	p.Scale(2)
	fmt.Printf("after factor: %v \n",p)

	p1 := &Point{2,4}  //这里取指针
	fmt.Printf("befor p1:%v \n", p1)
	//这个操作不会产生任何影响,因为这里操作的不是指针
	p1.Scale2(3)
	fmt.Printf("after p1:%v \n", p1)

}

//IntList 链表求和
type IntList struct {
	Value int
	Tail *IntList
}
//Sum 链表求和方法
func (list *IntList) Sum() int {	
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}



/// 测试结构体嵌套

//A4T 测试4
func A4T(){
	
}