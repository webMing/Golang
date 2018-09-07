package main

import (
	"fmt"
	"strconv"
	"bytes"
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

	fmt.Printf("buf: %v \n",intsToString([]int{3,4,5,6}))
	
 	//数 字符 转换
	num1 := 2334344
	nums1 := fmt.Sprintf("%d",num1) //method 1
	nums2 := strconv.Itoa(num1) //method 2
	
	fmt.Printf("%[1]T %[1]d \t %[2]T %[2]s \t %[3]T %[3]s \n",num1,nums1,nums2)
	
	//数值转换
	fmt.Printf("%[1]T \t %[1]s \n",strconv.FormatInt(15,16))	
	
	x, err := strconv.Atoi("12323")
	if err == nil {
		fmt.Println(x)
	}else {
		fmt.Println(err)
	}
     
        y, err := strconv.ParseInt("1010111",2,64)
	if err == nil {
		fmt.Println(y)
	}else {
		fmt.Println(err)
	}

}


func intsToString(valuse []int) string {

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i,v := range valuse {
		if i > 0 {
		    buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteByte(']')
	return buf.String()

}
