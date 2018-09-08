package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {

	fmt.Println("encoding/base64")
	message := "Away from keyboard https://www.golang.com!"
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeMessage)
	data, err := base64.StdEncoding.DecodeString(encodeMessage)
	if err == nil {
		fmt.Printf("data:%v \n", string(data))
	} else {
		fmt.Printf("err:%v \n", err)
	}

	//字符串拼接
	var buf = bytes.Buffer{} //声明为空对象
	buf.WriteRune('你')
	buf.WriteString("好啊，朋友")
	fmt.Println(buf.String())

}
