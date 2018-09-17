package  ch2  

import (
	"fmt"
)

//交换两个变量中的内容
func swap2(a, b *int) {
	*a,*b = *b,*a 
}

func swapValue2() {
	a,b := 1,2
	swap2(&a,&b)
	fmt.Printf("a:%d,b:%d \n",a,b)
}