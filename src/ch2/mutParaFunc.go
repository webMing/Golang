package main

import (
	"fmt"
)

func main() {
	custPrintf(1, 3, 4)
	fmt.Println("")
	custPrintf([]int{4, 5, 6, 7}...)
}

func custPrintf(s ...int) {
	for _, v := range s {
		fmt.Printf("v:%d \n", v)
	}
}
