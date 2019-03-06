//Package ch7 主要介绍 map的使用方法
package ch7

import (
	"bufio"
	"fmt"
	"os"
)

//MyMap 一些关于map的测试
func MyMap() {
	mapDefine()
}
func mapDefine() {
	var m0 map[string]int
	if m0 != nil {
		fmt.Println("map not nil")
		m0["lixao"] = 29
	}else {
		fmt.Println("map is nil")
	}

	//如果map为nil,那么 map["anyThing"] will result crash

	m1 := map[string]int{}
	if m1 != nil {
		fmt.Println("m1",m1)
	}
	// fmt.Println(m0)

	// var m2 map[string]int
	//如果map 为nil；那么对nil map 设置value 会崩溃
	// m2["wangxiaoer"] = 12
	// fmt.Println("m2",m2)

	// 利用 make 新建一个 map
	m3 := make(map[string]int)
	fmt.Println("m3 len:",len(m3))

	m4 := map[string]int{"lixiiao":29,"libei":30}
	fmt.Println("lenth of m4 is:",len(m4),"content of m4 is:",m4)

	//map  增，删，改，查
	//增
	m4["wangwu"] = 45
	fmt.Println("增加:content of m4:",m4)
	//改
	m4["wangwu"] = 12
	fmt.Println("改: content of m4:",m4)
	//查
	if v,ok := m4["testValue"]; !ok {
		fmt.Println("value of m4[\"testValue\"] not exist!")
	}else {
		fmt.Println("value of m4[\"testValue\"] is ",v)
	}
	//删
	delete(m4,"wangwu")
	fmt.Println("content of m4 is:",m4)

	//比较
	m5 := map[string]int{	
							"q":1,
							"a":2,
							"z":3,
						}

	m6 := map[string]int{
							"q":1,
							"a":2,
							"z":3,
							"x":4,
						}
	if compare(m5,m6) {
		fmt.Println("m5 == m6")
	}else {
		fmt.Println("m5 != m6")
	}

	// 把 map 当做 set 来使用
	dis := map[string]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := scanner.Text()
		if dis[line] == false {
			dis[line] = true
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error happen")
		os.Exit(1)	
	} 

}

func compare(m1,m2 map[string]int)(equal bool) {
	equal = true
	if len(m1) != len(m2) {
		equal = false
	}
	for k,m1V := range m1 {
		if m2V,ok :=  m2[k]; !ok || m1V != m2V{
			equal = false
			break
		}
	}
	return equal;

}