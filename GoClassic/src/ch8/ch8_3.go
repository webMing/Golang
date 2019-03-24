package ch8

import (
	"time"
	"fmt"
	"image/color"
)

//ColoredPoint 自定义的类型
type ColoredPoint struct {
	Point
	color color.RGBA
}

//A5T 测试5
func A5T() {
	var cp1 ColoredPoint
	cp1.X = 3
	cp1.Point.Y = 4
	fmt.Printf("cp1 Vlaues :%v\n", cp1)

	var cp2 ColoredPoint
	cp2.Point = Point{3,7}
	fmt.Printf("cp2 Vlaues :%v\n", cp2)
	cp2.Scale(2)
	fmt.Printf("cp2 after factor Vlaues :%v\n", cp2)

	fmt.Printf("cp1 and cp2 distance :%.2f \n", cp1.Distance(cp2.Point))
    
}

//下面这些内容关于方法变量

//A6T 测试6
func A6T() {
	p1 := Point{1,2}
	distanceFromP1 := p1.Distance
	p2 := Point{1,3}
	fmt.Printf("distance of p2 :%.2f \n", distanceFromP1(p2))
	fmt.Printf("var method type :%T \n",distanceFromP1)
	sacle := p1.Scale
	fmt.Printf("var method type :%T \n",sacle)
}


type rocket struct {
	content string	
}

func (r *rocket) lanch() {
	fmt.Println("Rock Fire!")
}
//A7T 测试7
func A7T() {
	rocket := rocket{"Rocket fire!"}
	time.AfterFunc(3 * time.Second, rocket.lanch)
	time.Sleep(5 * time.Second)
}

//A8T 测试8 方法表达
func A8T() {
	p1 := Point{1,4}	
	distanceFromP1 := p1.Distance
	p2 := Point{1,8}
	fmt.Printf("Distance form p2 %.1f \n",distanceFromP1(p2))
	
	distacne := Point.Distance
    fmt.Printf("Distance bett: %.1f \n",distacne(p1,p2))
	fmt.Printf("Distance Type: %T \n",distacne)
	
	scale := (*Point).Scale
	fmt.Printf("scale Type: %T \n",scale)
	scale(&p1,2) 
    fmt.Printf("sacle bett: %v \n",p1)
	
}

//Queue 自定义的一个队列
type Queue []int 

func (q *Queue) push(v int) {
	*q = append(*q, v)
}

//A9T 测试9 slice 是否会改变
func A9T() {

	s := []int{1,2,3,4,5}
	chageSliceValue(s)	
	fmt.Printf("values %v \n", s)

	fmt.Println("--------------------------------------")
	q := make(Queue, 1)
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	fmt.Printf("Queue type :%v \n",q)
	
	fmt.Println("--------------------------------------")
	f  := []int{}
	push(&f,1)	
	push(&f,2)	
	push(&f,3)	
	fmt.Printf("content of f:%v \n", f)

	fmt.Println("--------------------------------------")
	m := []int{1}
	m = append(m, 2)
	m = append(m, 3)
	m = append(m, 4)
	fmt.Printf("contentof m :%v \n", m)

	fmt.Println("--------------------------------------")
	mon := [...]string{"","1","2","3","4","5","6","7","8","9","10","11","12"}
	fmt.Printf("len:%d cap:%d \n",len(mon),cap(mon))
	q1 := mon[0:len(mon) - 1]
	fmt.Printf("q1:%v  %d\n", q1,len(q1))

}

func push(q *[]int, v int) {
	*q = append(*q, v)		
}

func chageSliceValue(s []int) {
	for i := range s {
		s[i] = 0
	}
}