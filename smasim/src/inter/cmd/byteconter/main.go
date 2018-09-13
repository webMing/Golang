package main

import (
	"io"
	"fmt"
	"inter"
)

func main() {
	byteConter()
}

func byteConter() {
	//直接调用方法(注意在不同的包中,变量的引用要添加包名)
	var bc inter.ByteConter
    n, err := bc.Write([]byte{1,2,3,4,5}) 
    if err != nil {
		fmt.Printf("the error is %v \n",err)	
	}else {
		fmt.Printf("conter:%v\n",n)
	}

	//interface
	bc = 0
	fmt.Fprintf(&bc,"%s","hello")
	fmt.Printf("conter:%d\n",bc)

	var w io.Writer
	w = &bc  //Write是ByteConter指针的方法所以添加&
	fmt.Printf("type of bc : %T",w)

}