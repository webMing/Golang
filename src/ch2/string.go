package main

import (
	"fmt"
	"strings"
	"bytes"
	)

func main() {
	fmt.Println("hello world")
	stracer := "死神来了,神 bye bye 爱不!"
        b_b := strings.Index(stracer,"bye")	
        b_e := strings.Index(stracer,"爱")	
	fmt.Printf("%s \n",stracer[b_b:b_e])

	//字符串链接
	fmt.Println("A" + "B")
	fmt.Println(strings.Join([]string{"A","B"},"&"))
	
	var buf bytes.Buffer
	a_s := "Hello"
	b_s := "shu!"
	buf.WriteString(a_s)
	buf.WriteString(b_s)
  	fmt.Println(buf.String())	
}
