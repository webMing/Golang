package ch2

import (
	"fmt"
	"flag"
)

func FlagParse() {
	var mode = flag.Int("mode",0,"参数说明")
	flag.Parse()
	fmt.Println("命令行参数：%d \n",*mode)
}