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
	"strings"
	"net/url"
	"log"
	"fmt"
	"net/http"
)

// 读取header
func header(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	fmt.Println(r.URL.Host)
	fmt.Println("Header Content")
	// for k, v := range r.Header {
	// 	fmt.Println(k,v)
	// }
	fmt.Println(r.Header.Get("Accept"))
	fmt.Fprintln(w, r.Header)
}

// 处理所有请求
func all(w http.ResponseWriter, r *http.Request) {
	// 处理所有的请求
	r.ParseForm()
	paths := strings.Split(r.URL.Path,`/`)
	for _,subPath := range paths {
		fmt.Println(subPath)
	}
	fmt.Println(r)
}

// 读取body
func body(w http.ResponseWriter,r *http.Request) {
	// r.ParseForm()
	// fmt.Println(r.FormValue("age"))
	//err := r.ParseMultipartForm(2024)
	//if  err != nil  {
		//og.Fatal(err)
	//}
	/*
	fmt.Printf("ContentLenth:%d \n",r.ContentLength)
	body := make([]byte, r.ContentLength)
	// body := []byte{}
	n,err := r.Body.Read(body)
	if n == 0 && (err != nil) {
		log.Fatalln(err)
	}
	fmt.Printf("Body Content:%s", body)
	fmt.Fprintf(w,"%s",body)
	*/
	fmt.Println(r)
}

// Extern6 把request的请求读到 terminal中
func Extern6() {
	server := http.Server {
	   Addr:"localhost:8090",
	}
	http.HandleFunc("/", all)
	http.HandleFunc("/header",header)
	http.HandleFunc("/body",body)
	server.ListenAndServe()
}

// Extern7 理解net/URL 包中的例子
func Extern7() {
	//u, err := url.Parse("http://www.bai.com:9090/search/to/book?q=donet")

	u, err := url.Parse("http:www.bai.com:9090/search/to/book?q=donet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	//u.Host = "google.com"
	//q := u.Query()
	//q.Set("q", "golang")
	//u.RawQuery = q.Encode()
	fmt.Println(u)
}

// Extern8 Roundtrip 官方所谓的RoundTrip
func Extern8() {
	// Parse + String preserve the original encoding...
	u, err := url.Parse("https://example.com/foo/%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())

}