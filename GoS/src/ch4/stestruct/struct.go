package stestruct

import (
	"fmt"
	//"fmt"
	"time"
)

type Employee struct {
	Name string
	ID   string 
	Address string 
	Dob		time.Time 
	Postion  string

}

type Point  struct{
	 X,Y float64
}

//在没判断struct 为nil 

//type unit int 

func EmployeeTest() {
   
	a := "abc"
	fmt.Printf("the a value is : %p \n",&a)
	 fmt.Println(a)
	 fmt.Printf("type:%d",len("hellp"))

	 for _,v :=  range "abc" {
		 fmt.Println(v)
	 }

}


func pt(x,y float64) Point {
	return  Point{x,y}
}