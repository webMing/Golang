package ch8

import (
	"log"
	"fmt"
	"time"
)

//Q4 关于时间一些内容
func Q4() {
	displayCurrentTime()
	// formatTime()
	// initTime()
	// cuParaseUnixTime()
}

// 在不同的时区显示当前的时间
func displayCurrentTime() {
	ltime := time.Now()
	local, err := time.LoadLocation("Asia/Tokyo")
	if err != nil { log.Fatal(err)
	}
	fmt.Println(ltime.In(local))
}

// 时间格式化
func formatTime() {
	format := "2006-01-02 15:04:05 Mon"
	currentTime := time.Now()
	local, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentTime.In(local).Format(format))
}

// 时间初始化
func initTime() {
   local, err := time.LoadLocation("Local")	
   if err != nil {
	   log.Fatal(err)
   }
   diplayTime := "2016-11-28 19:36:25"
   format := "2006-01-02 15:04:05"
   //cuTime, err := time.Parse(format, diplayTime)	
   cuTime, err := time.ParseInLocation(format, diplayTime, local)
   if err != nil {
	  	log.Fatal(err)
   }
   fmt.Println(cuTime)
}

// 解析UnixTime
func cuParaseUnixTime() {
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	sec := int64(1480390585)
	parsedTime := time.Unix(sec,0)
	fmt.Println(parsedTime.In(location))
}
