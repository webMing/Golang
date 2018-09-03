package main

import (
	"fmt"
	)

type Score []float32

func main() {
	fmt.Println("-----------------------")
	s := Score{11.0,23,45,64}	
	fmt.Printf("s: %v \n",s)	
        s.changeValue(0,0)	
	fmt.Printf("s: %v \n",s)	
	fmt.Println("-----------------------")
}

func (s Score) changeValue(i int, value float32) {
	if i > len(s) -1 {
		return 
	}
        s[i] = value
} 
