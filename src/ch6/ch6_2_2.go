package main

import (
	"fmt"
	)

type Value map[string][]string


func (v Value) Get(key string) string {
	if v := v[key]; len(v) != 0 {
		return v[0] 
	}
	return ""
}

func (v Value) Add(key, value string) {
	v[key] = append(v[key],value)
}


func main() {
      name := Value{"name":{"li"}}
      fmt.Printf("names : %v \n",name.Get("name"))
      fmt.Printf("names : %v \n",name.Get("score"))
      name.Add("wang","haoaho")
      fmt.Printf("names : %v \n",name.Get("wang"))
	
      var books Value
      books.Add("en","swift")
}
