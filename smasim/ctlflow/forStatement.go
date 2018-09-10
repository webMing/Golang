package ctlflow

import (
	"fmt"
)

// ForStatement 介绍for的几种用法
func ForStatement() {
	//for的普通刑事
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("use case one : %d \n",i)
	}

	var i = 0
	for ; i < 10 ; i++ {
		fmt.Printf("use case two :%d \n",i)
	}

	var j = 0
	for ; j < 10 ; {
		fmt.Printf("use case three:%d\n",j)	
		j++
	}

	var k = 0
	for ; ; k++ {
		if k > 10 {
			break
		}
		fmt.Printf("use case four:%d\n",k)
	}

	 var f = 0
	 for {
		if f < 10 {
			fmt.Printf("use case five:%d\n",f)
			break
		}
		f++
	 }

}