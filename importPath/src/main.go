package main

import (
	 "fmt"
	 "github.com/ste/stringutil"
	 "gopl.io/ch2/tempconv"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(stringutil.Reverse("123456"))
	n := tempconv.BoilingC
	fmt.Println(n)
}