//
// ch11_5.go
// ch11_5
//
// Descrition
//
// Created by Stephanie on 2019/04/24.
// Copyright © 2019 Stephanie. All rights reserved.
//

package ch11

import (
	"log"
	"fmt"
	"net/http"
)

// 读取header
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Header Content")
	// for k, v := range r.Header {
	// 	fmt.Println(k,v)
	// }
	fmt.Println(r.Header.Get("Accept"))
	fmt.Fprintln(w, r.Header)
}

// 读取body
func body(w http.ResponseWriter,r *http.Request) {
	fmt.Printf("ContentLenth:%d \n",r.ContentLength)
	body := make([]byte, r.ContentLength)
	n,err := r.Body.Read(body)
	if n == 0 && (err != nil) {
		log.Fatalln(err)
	}
	fmt.Fprintf(w,"%s",body)
}

// Extern6 把request的请求读到 terminal中
func Extern6() {
	server := http.Server {
	   Addr:"localhost:8090",
	}
	http.HandleFunc("/header",header)
	http.HandleFunc("/body",body)
	server.ListenAndServe()
}