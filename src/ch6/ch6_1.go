package main

import (
	"fmt"
	"math"
	)

type Point struct{X, Y float64}
type Path []Point

func main() {
          c := Point{2, 3}
	  d := Point{2, 4}
	  fmt.Printf("distance %v and %v is : %.2f \n",c,d,Distance(c,d))
	  fmt.Printf("distance %v and %v is : %.2f \n",c,d,c.Distance(d))
		
	  path := Path{{1,1},{5,1},{5,4},{1,1}} 
	  fmt.Printf("distance of path is :%.2f \n",path.Distance())

}

func Distance(a,b Point) float64 {
	return math.Hypot(a.X - b.X,a.Y - b.Y)
}

func (a Point) Distance(b Point) float64 {
	return Distance(a,b)
}

func PathDistance(p Path) float64 {
	sum := 0.0  
	for i := range p {
		if i > 0 {
			sum += Distance(p[i-1],p[i])
		} 
	}
	return sum 
}

func (p Path) Distance() float64 {
	return PathDistance(p)
}
