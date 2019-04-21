//
// ch11_3.go
// ch11_3
//
// Descrition
//
// Created by Stephanie on 2019/04/15.
// Copyright © 2019 Stephanie. All rights reserved.
//

package ch11

import (
	// "html/template"
	"fmt"
	"net/http"
)

// Extern4 可导出的函数
func Extern4() {
	mux :=  http.NewServeMux()
	files := http.FileServer(http.Dir("/Users/stephanie/tmp"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	// mux.HandleFunc("/", index)
	f := func(w http.ResponseWriter, r *http.Request) {
		/*
			Scheme     string
			Opaque     string    // encoded opaque data
			User       *Userinfo // username and password information
			Host       string    // host or host:port
			Path       string    // path (relative paths may omit leading slash)
			RawPath    string    // encoded path hint (see EscapedPath method)
			ForceQuery bool      // append a query ('?') even if RawQuery is empty
			RawQuery   string    // encoded query values, without '?'
			Fragment   string  
		*/
		url := r.URL
		fmt.Fprintf(w, "URL Scheme:%q  \n", url.Scheme)	
		fmt.Fprintf(w, "URL Opaque:%q  \n", url.Opaque)	
		fmt.Fprintf(w, "URL UserInfo:%v \n", url.User)	
		fmt.Fprintf(w, "URL Host:%q \n", url.Host)	
		fmt.Fprintf(w, "URL Path:%q \n", url.Path)	
		fmt.Fprintf(w, "URL RawPath:%q \n", url.RawPath)	
		fmt.Fprintf(w, "URL ForceQuery:%v \n", url.ForceQuery)	
		fmt.Fprintf(w, "URL RawQuery:%q \n", url.RawQuery)	
		fmt.Fprintf(w, "URL Fragment:%q \n", url.Fragment)	

	}
	mux.HandleFunc("/", f)
	s := &http.Server {
		Addr:":8090",
		Handler:mux,
	}
	s.ListenAndServe()	
}

/*
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"template/layout.html",
		"template/navbar.html",
		"template/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.threads(); if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
*/