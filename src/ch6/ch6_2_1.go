package main

import (
	"fmt"
	)

type IntList struct {
	Value int;
	Tail *IntList  
}

func main(){

	//create data
	first := &IntList{0,nil}

	last := first
	for _,v := range []int{1,2,3,4,5,6} {
		last = last.Append(v)
	}	

	
        //打印
	for begin := first; begin != nil ; begin = begin.Tail {
		fmt.Printf("The value is: %d \n",begin.Value)
	} 

}


func (li *IntList) Append(v int) *IntList {
	s := &IntList{v,nil}
	li.Tail = s
	return s
} 

