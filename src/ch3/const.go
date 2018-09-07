package main

import (
	"fmt"
	"math"
	"unsafe"
	)



const (
	a = 1
	b  // = 1 
	c = 4
	d  // = 4
	)


const (
	_ = iota
	X
	Y	
)

const (
	_ = 1 << (10 * iota)
	KiB 
	MiB
	GiB 
	)
func main() {
	fmt.Printf("\n")
	fmt.Printf("Const a = %d, b = %d, c = %d, d = %d \n",a,b,c,d)
	fmt.Printf("Const X = %d,Y = %d \n",X,Y)
	fmt.Printf("Const KiB = %v,MiB = %d \n",KiB,MiB)
	fmt.Printf("Const matho.Pi = %g\n",math.Pi)

	fmt.Printf("%T sizeOf(int):%d \n",0,unsafe.Sizeof(0))
	fmt.Printf("%T \n",0.0)
	fmt.Printf("%T \n",0i)
	fmt.Printf("%T \n",'\000')

}

