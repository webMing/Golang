package ch9

import (
	"os"
	// "gopl.io/ch5/links"
	// "html"
	"fmt"
)

// Q7 测试
func Q7() {
	allocL()
} 

type person struct{
	name string
	age int
}
var a = 32

func allocL() {
	fmt.Printf("%p \n",&a)

	p := new(person)
	println(p)
	fmt.Printf("%p \n",p)
	fmt.Printf("%p \n",&p)

	c := [100]int{}
	fmt.Printf("%p \n",&c)


	str := "This is string!@!"
	fmt.Fprintf(os.Stdin, "%s", str)
	var b []byte
	os.Stdin.Read(b)
	fmt.Printf(" =    = =======  %s",b)

}