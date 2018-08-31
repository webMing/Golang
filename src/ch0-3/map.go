package main

import (
	"fmt"
	"sort"
	)

func main() {
	
	fmt.Println("================")

	var map_0 = make(map[string]int)
	map_0["one"] = 1
	map_0["two"] = 2
	fmt.Printf("map_0 : %v \n",map_0)	
	
	map_1 := map[string]int{"three":3,"four":4}
	fmt.Printf("map_1 : %v \n",map_1)

	ages := make(map[string]int)
  	ages["li"] = 23
	ages["zhao"] = 24
	for k,v := range ages {
		fmt.Printf("key: %s, value:%d \n",k,v)
	}
	
	//创建空字典
	//emptyMap := map[string]int{}
	var emptyMap map[string]int 
	if emptyMap == nil {
		fmt.Printf("emptyMap is nil! \n")
	}

	if len(emptyMap) == 0 {
		fmt.Printf("emptyMap is empty! \n")
	}
	
	delete(ages,"li")
	fmt.Printf("map ages:%v \n",ages) 
	
	ages["bob"] = ages["bob"] + 1
	ages["bob"] +=  1
	ages["bob"]++
	//fmt.Printf("take address:%v",&ages["bob"])
	fmt.Printf("ages[\"bob\"] : %v \n",ages)
	
	letters := map[string]string{"a":"A","c":"C","d":"D"}
	lettersKey := []string{}
	for k,_ := range letters {
		lettersKey = append(lettersKey,k)
	}
	sort.Strings(lettersKey) //排序
	for _,v := range lettersKey {
		fmt.Printf("letters[%s] : %v \n",v,letters[v])
	}

	fmt.Println("================")

}
