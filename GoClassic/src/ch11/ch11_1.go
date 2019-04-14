//
// ch11_1.go
// ch11_1
//
// Descrition
//
// Created by Stephanie on 2019/04/14.
// Copyright © 2019 Stephanie. All rights reserved.
//

package ch11

import (
	"fmt"
	"net/http"
)

// Extern1  go_web 1-1 codes
func Extern1() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!",request.URL.Path[1:])
}
