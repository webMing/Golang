package main

import (
	"fmt"
	//"ch5"
)

func main() {
	//ch5.LearnArray()
	s := make([]int,3)[3:]
	fmt.Println("len:",len(s),"cap(s):",cap(s))
}