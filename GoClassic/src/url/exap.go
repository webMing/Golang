//
// exap.go
// exap
//
// code display 
//
// Created by Stephanie on 2019/05/11.
// Copyright © 2019 Stephanie. All rights reserved.
//

package url

import (
	"fmt"
	"log"
	"net/url"
)

/**
	Scheme     string
	Opaque     string    // encoded opaque data
	User       *Userinfo // username and password information
	Host       string    // host or host:port
	Path       string    // path (relative paths may omit leading slash)
	RawPath    string    // encoded path hint (see EscapedPath method)
	ForceQuery bool      // append a query ('?') even if RawQuery is empty
	RawQuery   string    // encoded query values, without '?'
	Fragment   s
*/

//Exap1 一些例子
func Exap1() {

	// 当前解析格式为 scheme:opaque[?query][#fragment]
	u,err := url.Parse(`http:www.baidu.com:8080/path/to/source?q=index`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
}

//Exap2 一些例子
func Exap2(){

	//解析格式[scheme:][//[userinfo@]host][/]path[?query][#fragment]
	u,err := url.Parse(`https://www.baidu.com:8080/path/to/your/resource?index=2&page=3#23`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Port())
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.Query())
	fmt.Println(u.Fragment)
	
}


//Exap3 RoundRrip means ? url中包含斜线时候如何解析
func Exap3() {
	url,err := url.Parse(`https://www.baidu.com:8080/path/to/your%2fsearch?index=2&count=3#4`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url.Path)	
	fmt.Println(url.RawPath)
	fmt.Println(url.RequestURI())

	fmt.Println(url.EscapedPath())

}