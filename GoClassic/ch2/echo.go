package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	// echo2()
	// echo3()
	// processName()
	argsIndexAndValue()
	// fmt.Println("Usage of this!")
}

// echo1 输出命令行参数
func echo1() {
	var display,sep string
	for index := 1; index < len(os.Args); index++ {
		 display += sep + os.Args[index];
		 sep = " "
	}
	fmt.Println(display)
}

// echo2 输出命令行参数
func echo2() {
	var display,sep string
	for _,arg := range os.Args[1:] {
		display += sep + arg
		sep = " "
	}
	fmt.Println(display)
}

// echo3  输出命令行参数
func echo3() {
	fmt.Println(strings.Join(os.Args[1:]," "))	
}

// 输出程序名字
func processName() {
	fmt.Println(os.Args[0])
}

// 输出参数的索引和值,每行一个 
func argsIndexAndValue() {
	for index,value := range os.Args[1:] {
		fmt.Printf("index:%d value:%v \n",index, value)	
	}
	time.Now()	
}