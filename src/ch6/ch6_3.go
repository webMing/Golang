package main

import (
	"fmt"
	"image/color"
	"math"
	)

type Point struct{X,Y float64}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X - q.X , p.Y - q.Y)	
}

func (p *Point) ScaleBy(fac float64) {
	p.X *= fac
	p.Y *= fac
}

type ColorPoint struct{ 
			*Point
			Color color.RGBA
		      }


func main() {
	fmt.Println("works well")	

	var cp ColorPoint
	cp.Point = &Point{1,1} 
        fmt.Printf("----------cp value is %v \n",*cp.Point) 

	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}

	var p = ColorPoint{&Point{1,1},red}
	var q = ColorPoint{&Point{5,4},blue}
	fmt.Println(p.Distance((*q.Point)))
	
	p.ScaleBy(0.5)
	fmt.Printf("p value is %v \n",p)
}



