package main

import (
	"fmt"
	)


func main() {
	
	aStr := []string{"ni","hao"," "," ", "zhang","ya","long","!"}
	tempStr := ""
	for _,s := range aStr {
		tempStr = fmt.Sprintf("%s%s",tempStr,s)
	}  
	fmt.Printf("tempStr:%s \n",tempStr)	
     
        //切片	
	tempStrSlice := nonempty(aStr[:])
	
	tempStr = "" 
	for _,s := range tempStrSlice {
		tempStr = fmt.Sprintf("%s%s",tempStr,s)
	}  
	fmt.Printf("tempStr:%s\n",tempStr)
	
	//len("") != len(" ")
	//fmt.Printf("len(\"\") is :%d \n",len("")) //要进行转义
	//fmt.Printf("len(\" \")is :%d \n",len(" ")) //要进行转义
	
	numbersArray := [...]int{1,2,3,4,5,6}
	numbersSlice := numbersArray[:]
	fmt.Printf("numbersSlice is : %v\n",numbersSlice)
	numbersSlice = deleteElement(2,numbersSlice)
	fmt.Printf("numbersSlice is : %v\n",numbersSlice)
}

// 去掉空格...
func nonempty(strings []string) []string {
	i := 0
	for _,s := range strings {
		if s != " " {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//删除中间的某个元素
func deleteElement(index int, nums []int) []int {
	//通过copy命令将后面的数据往前面移动
	copy(nums[index:],nums[index + 1:])
	//copy(nums[index:],nums[(index + 1):])
	return nums[:len(nums) - 1]
}
