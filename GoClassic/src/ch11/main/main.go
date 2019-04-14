package main

import (
	"fmt"
	"ch11"
)

func main() {
	s := "/a/b/c/d!"
	sSlice := s[1:]
	fmt.Println(sSlice)
	
	// ch11.Extern1()	
	ch11.Extern2()
}