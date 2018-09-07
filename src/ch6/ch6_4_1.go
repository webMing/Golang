package main

import (
	"math"
	"fmt"
	)


type Point struct{X, Y float64}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X , p.Y - q.Y)
}

func (p *Point) ScaledBy(fac float64) {
	(*p).X *= fac
	(*p).Y *= fac 
}

func main() {
        var a = Point{1,1}
	distanceFromA := a.Distance
	s := distanceFromA(Point{1,2})
	fmt.Printf(" s : %g \n",s)	
	
	scaledP := a.ScaledBy
	scaledP(0.5)
	fmt.Printf("a :%g \n",a)
	
	//值表达式
	p := Point{4,4}
	(*Point).ScaledBy(&p,0.5)
	fmt.Printf("v: %v \n",p)
}
