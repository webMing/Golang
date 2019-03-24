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

