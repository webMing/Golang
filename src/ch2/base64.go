package main

import (
	"fmt"
	"encoding/base64"
	)

func main() {

	fmt.Println("encoding/base64")
	message := "Away from keyboard https://www.golang.com!"
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeMessage)
	data, err := base64.StdEncoding.DecodeString(encodeMessage)
	if err == nil {
		fmt.Printf("data:%v \n",string(data))
	}else {
		fmt.Printf("err:%v \n",err)
	}
}
