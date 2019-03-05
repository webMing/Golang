package main

import (
	"fmt"
	"ch6"
)

func main() {
	ch6.F1();
	ch6.F2("Stephanie");
	a := ch6.F3("li xiao")
	fmt.Println(a)
	ch6.F4(2,4,6,7)
	f := func(a int)(b int){
		return a + 4
	}
	ch6.F5(f,5)
}