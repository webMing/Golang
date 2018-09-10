package ctlflow

import (
	"fmt"
)

// ContentStatment   条件语句 
func ContentStatment() {
	//普通声明
	var  a = 10
	if a > 10 {
		fmt.Println("a > 10")
	}else if a < 10 {
		fmt.Println("a < 10")
	}else if a == 10 {
		fmt.Println("a == 10")
	}

	//特殊条件语句u; 注意a,b 的作用域
	if a, b := 10,40;  a > b {
		fmt.Println("a > b")
	}else if a < b {
		fmt.Println("a < b")
	}else if a == b {
		fmt.Println("a == b")
	}
 
	var c = 2
	switch c {
		case 1:{
			fmt.Println("c == 1")
		}
		case 2 :{
			fmt.Println("c == 2")
		}
		case 3 :{
			fmt.Println("c == 3")
		}
		default :{
			fmt.Println("unkown")
		}
	}

	var d = 3
	switch d {
		case 1,4 : {
			fmt.Println("1,4")
		}
		case 2,3:{
			fmt.Println("2,3")
		}
		case 5,6:{
			fmt.Println("5,6")
		}
	}

	var m = 5
	switch {
		case m == 4 :{
			fmt.Println("m == 4")
		}
		case m == 3 :{
			fmt.Println("m == 3")
		}
		case m == 5 :{
			fmt.Println("m == 5")
		}
	}
}