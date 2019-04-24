//
// ch11_4.go
// ch11_4
//
// Descrition
//
// Created by Stephanie on 2019/04/24.
// Copyright © 2019 Stephanie. All rights reserved.
//

package ch11

import (
	"time"
	"fmt"
	"net/http"
)

// MyHandler 用来处理所有的链接
type MyHandler struct{}
func (h MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "Time:%s", time.Now().Format("2006-01-02 15:04:05"))
}

// Extern5 测试处理函数 
func Extern5() {

	fmt.Print("创建服务器 \n")
	/*
	rootFun := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Root:%s Time:%s",r.URL,time.Now().Format("2006-01-02 15:04:05"))
	}
	aFun := func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s",r.URL)	
	}
	bFun := func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL:%s Current Time :%s",r.URL,time.Now().Format("2006/01/02 03:04:05"))	
	}
	http.HandleFunc("/",rootFun)
	http.HandleFunc("/a/",aFun)
	http.HandleFunc("/b/",bFun)
	*/

	handler := MyHandler{}
	server := http.Server {
		Addr:"localhost:8090",
		// Handler:handler,
	}
	http.Handle("/root/", handler)
	server.ListenAndServe()
}
