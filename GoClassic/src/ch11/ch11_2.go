//
// ch11_2.go
// ch11_2
//
// Descrition
//
// Created by Stephanie on 2019/04/14.
// Copyright © 2019 Stephanie. All rights reserved.
//

package ch11

import (
	"gopl.io/ch4/github"
	"html/template"
	"log"
	"time"
	"os"
)

const templ = ` {{.TotalCount}} issues:
{{range .Items}} -----------------------
Number:{{.Number}}
User:{{.User.Login}}
Title:{{.Title|printf ".%64s" }}
Age:{{.CreatedAt | daysAgo}} days
{{end}}`

var t,err = template.New("issueList").Funcs(template.FuncMap{"daysAgo":daysAgo}).Parse(templ)
var report = template.Must(t, err)
// Extern2 go 圣经4.6
func Extern2() {
	// result,err := github.SearchIssues(os.Args[1:]) 
	result,err := github.SearchIssues([]string{"repo:golang/go is:open json decoder"})
	check(err)
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time)int {
	return int(time.Since(t)/24)
}

// Extern3 go example 
func Extern3() {
	const tpl = `
	<!DOCTYPE html>
	<html>
 		   <head>
    		    <meta charset="UTF-8">
     		   <title>{{.Title}}</title>
    		</head>
    	<body>
			{{range .Items}}
				<div>
					{{ . }}
				</div>
			{{else}}
				<div>
					<strong>no rows</strong>
				</div>
			{{end}}
    	</body>
	</html>`

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	/*
	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)
	*/	

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
