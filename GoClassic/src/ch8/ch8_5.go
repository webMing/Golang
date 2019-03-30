package ch8

import (
	"fmt"
	"time"
)

// Q3 导出的方法
func Q3() {
	go display5(100 *time.Millisecond)	
	//time.Sleep(time.Hour)
	const v = 45
	fmt.Printf("Result:%d",flib(v))
}

func flib(v int) (n int) {
	if v < 2 {
		return v 
	}
	return flib(v - 1) + flib(v - 2)
}
func display5(delay time.Duration) {
	for {
		for _, r := range `-/|\`{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}	
}

