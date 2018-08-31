package main

import (
	"fmt"
	) 

func main() {
	fmt.Printf("===============================\n")
	//切片生成
	//var a []int = []int{1,3,4}
	//var a = []int{1,3}
	//var a = []int{100:0}
	//a := []int{1,3}
	test := [...]int{34:3}
	var a  = test[1:10] //不包括最后一个下标
	fmt.Printf("len :%d,cap : %d \n",len(a),cap(a))	


	var tq byte = 4 //byte uint8
 	fmt.Printf("%[1]T %08[1]b \n",tq)	

	months := [...]string{1:"January",2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",8:"August",9:"Septmeber",10:"October",11:"November",12:"December"}

	for index,value := range months{
		fmt.Printf("months[%d] = %s \n",index,value)
	}

	Q2 := months[4:7]
	for index,value := range Q2{
		fmt.Printf("Q2[%d] = %s \n",index,value)
	}	
	
	summer := months[4:7]
	for index,value := range summer {
		fmt.Printf("summer[%d] = %s \n",index,value)
	}	
	
	//fmt.Println(summer[3]) 超出范围将产生异常
	
	//字符串切片
	testStr := "abcdefghjk"
	testRune := testStr[:]
	for index,value := range testRune {
		fmt.Printf("testRune[%d] = %q \n",index,value)
	}

	//var testHex uint32 = 0xf0f1f2f3
	//testHexSlice := testHex[:]
	//fmt.Println(testHexSlice)

	var runes []rune
	for _,value := range "hello 石家街" {
		 runes = append(runes,value)
	}
	for i,v := range runes {
		fmt.Printf("runes[%d] = %q \n",i,v)
	}
	
	var s []int
	fmt.Printf("len(s) = %d ;cap(s) = %d \n",len(s),cap(s))
	fmt.Printf("test is nil:%v \n",[]int(nil))

	//append
	aa := []int{}
	fmt.Printf("aa: %v\n",aa)
	aa = append(aa,4,5,6)
	aa = append(aa,3,4,5)
	bb := []int{44,66}
	aa = append(aa,bb...)
	fmt.Printf("aa: %v\n",aa)
	bb[0] = 99
	bb[1] = 99 //改变之后是值？ 还是指针？
	bb = append(bb,99)
	fmt.Printf("aa: %v \n",aa)
	fmt.Printf("===============================\n")
	
}


func addpendInt(x [] int, y int) []int {

	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
	
		z = make([]int,zlen,zcap)
		copy(z,x)
  	}

	z[len(x)] = y
	return z
}
