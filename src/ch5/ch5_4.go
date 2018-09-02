package main

import (
	"fmt"
	)

func main() {
	
	fmt.Println("------------------------")

	//分开形式
	m := map[string][]string{}
	m["name"] = []string{"li","wang","zhao"}
	m["score"] = []string{"29","34","56"}
	fmt.Printf("m value : \n%v\n",m)

	//组合形式(值的字面量)	
	m1 := map[string][]string {
		"name" : {"李","赵","王"},
		"score" : {"34","45"},
	}
	fmt.Printf("m1 value : \n%v\n",m1)

	fmt.Println("------------------------")

}
