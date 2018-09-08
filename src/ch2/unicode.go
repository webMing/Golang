package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//如果是英文使用效果是相同的.
	fmt.Println(len("abcdef"))
	fmt.Println(utf8.RuneCountInString("abcdef"))
	fmt.Println(utf8.RuneCountInString("你好!"))

	fmt.Println("\n")

	profile := struct {
		Name string
		Hp   int
	}{
		Name: "rat",
		Hp:   33,
	}

	fmt.Printf("%%v :\t  %v\n", profile)
	fmt.Printf("%%#v:\t  %#v \n", profile)
	fmt.Printf("%%+v:\t  %+v \n", profile)
}
