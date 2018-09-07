package main

import (
	"fmt"
	"sort"
	)

func main() {
	fmt.Println("===============")
	
	//有序打印	
	scores := map[string]float32{
					"zhao":23.3,
					"qian":34.4,
					"sun":12.4,
					"li":100.0 								    }
	names := make([]string,0,len(scores))
        for key := range scores {
		names = append(names,key)
	}
	sort.Strings(names)
	for _,key := range names {
		fmt.Printf("scores[%v]:%5.1f \n",key,scores[key])
	}	
	//fmt.Printf("scores:%v \n",scores)

	//辨别key是否存在map中
	carNum := map[string]string{"li":"1010","zhao":"1100","wang":"0011"}
	carUser := "li"

	//num,ok := carNum[carUser]
	if num,ok := carNum[carUser]; ok {
	    fmt.Printf("carNum of %v is :%v \n",carUser,num)
	}else {
	    fmt.Printf("carNum of %v is not eist\n",carUser) //注意函数的作用域
	}

	//深拷贝浅拷贝
	books := map[string]int{"li":1,"wang":2,"zhao":1}
	changeValue(books)
	fmt.Println(books)	
	
	//Slice 函数传值
	systems := []int{34,50,90}
	changeSlice(systems)
	fmt.Printf("the system is %v \n",systems)
	fmt.Println("===============")
}

func equal(x,y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k,v := range x {
		if value,ok := y[k]; !ok || value != v { //注意这里的书写方式
		   return false	
		}
	}
	return true 
}


//深拷贝还是浅拷贝
func changeValue(a map[string]int) {
	//经过测试试指针传递
	for k := range a {
		a[k] = 3 //set all value is 3
	}	
}

//Slice 通过函数改变其内容
func changeSlice(s []int) {
	for i := range s {
		s[i] = 1
	}
}
