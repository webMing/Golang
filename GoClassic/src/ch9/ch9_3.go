/**
*  Go圣经中转换为JSON格式的shuju 
 */

package ch9

import (
	"os"
	"encoding/json"
	"fmt"
	// "fmt"
	// "json"
	"log"
)

// Q3 测试JSON相关的一些内容
func Q3() {
	/*
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n",data)
	*/
	// 其他的实现
    var data []byte
	err := json.NewEncoder(os.Stdout).Encode(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}

// Movie 电影数据结构
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "cas", Year: 1934, Actors: []string{"A", "B"}},
	{Title: "孙悟空大闹天空", Year: 1200, Color: true, Actors: []string{"李肖", "王杰"}},
}
