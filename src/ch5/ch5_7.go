package main

import (
	"fmt"
	)

func main() {

	fmt.Printf("----------------\n")
	// 可变参数
        fmt.Printf("Sum :%d \n",cSum(1,2,3,4,5))	
        fmt.Printf("Sum :%d \n",cSum(3,4,5))	
	score := [3]int{1,1,1}
        fmt.Printf("Sum :%d \n",cSum(score[:]...))

	fmt.Printf("----------------\n")
}

func cSum(p...int) (res int) {
    for _,v := range p {
	res += v
    }
    return 
}
