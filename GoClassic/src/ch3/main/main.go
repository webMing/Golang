package main

import (
	"unsafe"
	"fmt"
	"ch3"
)

func main () {

	fmt.Println("Test!")
	ch3.BaseType()

	var u uint = 255
	fmt.Println(unsafe.Sizeof(u))
	ru := u + 1
	fmt.Printf(" ===============   %T \n",ru)
	fmt.Println(u,u+1,u*u)

	var i int8 = 127
	fmt.Println(i,i+1,i * i)

	// assic := 'c'
	unicode := '国'
	fmt.Printf("%[1]c %[2]T %[1]d %[1]c %[1]q",22269,unicode)

	var f float32 = 16777216
	fmt.Println(f == f + 1)
}