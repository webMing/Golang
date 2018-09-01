package main

import (
	"fmt"
	"time"
	"unsafe"
	)

func main() {

	fmt.Println("========================")
	type Employee struct {
		ID        int
		Name      string
		Address	  string
		DoB	  time.Time
		Positon   string
		Salary    int
		ManagerID int
	}
	var dilbert Employee
	//打印默认值
	fmt.Printf("ID :%d Name:%s Dob: %v \n",dilbert.ID,dilbert.Name,dilbert.DoB)
	
	type Student struct {
		Name 	string
		Id      string
		Score 	float32
	}
	//var yaLong Student
	var yaLong Student
	fmt.Printf("Name : %s, Id : %s , Score : %f \n",yaLong.Name,yaLong.Id,yaLong.Score)
	
	yaLong.Id = "123232434"
	fmt.Println("yaLong.Id",yaLong.Id)
        
	score := &yaLong.Score
	*score = 33.32 + *score //修改内容
	fmt.Println("score :",yaLong.Score)
	
	fmt.Printf("struct{} size is : %d \n",unsafe.Sizeof(struct{}{}))
	
	struct_map := make(map[string]struct{})
	//struct_map["a"] = struct{}{}
	if v,ok := struct_map["a"] ; ok {
		fmt.Printf("exit : %v \n", v)
	}
	fmt.Println("========================")

}
