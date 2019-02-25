package ch3 

import (
	"fmt"
	"math"
	"unsafe"
)

var localChar = '国'
const localVar = 4

//BaseType is go...
func BaseType() {
	if  math.NaN() != math.NaN(){
		fmt.Println("NaN != NaN")
	} 
	fmt.Printf("%T \n",localVar)
	fmt.Printf("%T \n",localChar)
	var a int64 = 4;
	b := 4.4
	fmt.Printf("a type:%T \n",a)
	fmt.Printf("b type:%T \n",b)
    
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(new(int8)))
	fmt.Println(unsafe.Sizeof(new(int64)))
	fmt.Println(unsafe.Sizeof(new(float64)))
}
