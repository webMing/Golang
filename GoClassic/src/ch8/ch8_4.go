package ch8

import (
	"fmt"
)

//A10T 测试10
func A10T() {
				
}

// Q1 类型推到
func Q1(x interface{}) (s string) {
	if x == nil {
		return "NULL"
	}else if _, ok := x.(int); ok {
		return fmt.Sprint(x)
	}else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d",x)
	}else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	}else if s, ok := x.(string); ok {
		 return s 
	}else {
		panic("unexpected type:%T value:%v \n,x,x")
	}

}

// Q2 类型推到2
func Q2(x interface{}) (s string) {
	 switch  x := x.(type) {
	 case nil:
			return "NULL"
	 case int:
		 return fmt.Sprintf("%d",x)
	 case uint:
		return fmt.Sprintf("%d",x)
	 case string:
		return x
	 case bool:
			if x {
				return "TRUE"
			}
			return "FALSE"
	 default:
		panic("unexpected Type:%T  Value:%v,x,x")
	 }
}