package main

import (
	"fmt"
	"bytes"
	"io"
//	"inter"
)

func main() {
	//inter.FlagValue()
	var w io.Writer
	var buf *bytes.Buffer
	w = buf         

	// 注意这里w 并不等于nil. 
	if w == nil {   
		fmt.Printf("w == nil!")
	}else {
		fmt.Printf("w != nil!")
	}

	
}