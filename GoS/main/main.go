package main

import (
	"fmt"
	//相对路径
	//"../ch2"
	"../inird"
	//"fmt"
)

func main() {
	//ch2.SwapValue1()
	//ch2.FlagParse()
	value, err := inird.GetValue("../inird/ini.txt", "Core", "hideDotFiles")
	if err != nil {
		fmt.Printf("%v \n", err)
	} else {
		fmt.Printf("%v \n", value)
	}
}