package main

import (
	"fmt"
	)

func main() {

	fmt.Printf("man func!\n")
	//strings slice
	s1 := "我的歌声里有☺😎"
	s2 := s1[:]
	fmt.Printf("%T  \t %T \n",s1,s2)
	
	/*	
	for _,r := range s2 {
		fmt.Printf("r : %c \n",r)
	}
	*/
	
	s3 := []rune{'我','☺','w'}
	fmt.Printf("%T \t %T \n",s3,string(s3))

	s4 := []byte{'a','c','d','1'}
	fmt.Printf("%T \t %T \n",s4,string(s4))
	fmt.Printf("%s \n",string(s4))
	
	s5 := "我的👓"
	for _,r := range []rune(s5) {
		fmt.Printf("%c \n",r)
	}


}
