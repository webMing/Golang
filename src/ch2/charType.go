package main

import (
        "unicode/utf8"	
	"fmt"
)

func main() {
	var a byte = 'a'
	b := '你'
        c := '\u4e14'
	fmt.Printf("type of a: %T \n",a) 
	fmt.Printf("type of b: %T \n",b) 
	fmt.Printf("type of c: %T \n",c)
	
	display()
}

func display() {

	s := "中国好声音，🌶了你笑脸22:-D😓险恶"
	fmt.Printf("the len of s is %d \n",len(s))
        fmt.Printf("the len of s is %d \n",utf8.RuneCountInString(s))	
	
	for i := 0 ; i < len(s) ; i++ {
	      r,size := utf8.DecodeRuneInString(s[i:])
	      fmt.Printf("%d \t %c \n",i,r)
	      i += size
	}

	fmt.Printf("\n\n\n")

	chars := s[:]
	for _,v := range chars {
		//循环次数就是字体个数
		fmt.Printf("%c \n",v)
	}
}
