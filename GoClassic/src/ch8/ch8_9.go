package ch8

import (
	"log"
	"fmt"
	"time"
)

//Q4 关于时间一些内容
func Q4() {
	ltime := time.Now()
	local, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ltime.In(local))
}