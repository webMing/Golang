package main

import (
	"flag"
	"fmt"
)

const usage = `
			   usage:
         	   1 .go run flag.go skillName
               2 .go run flag.go --skill=skillName
               `

var skillPara = flag.String("skill", "", usage)

func main() {

	flag.Parse()
	var skills = map[string]func(){

		"fly": func() {
			fmt.Println("fly skill")
		},

		"run": func() {
			fmt.Println("run skill")
		},

		"fire": func() {
			fmt.Println("fire skill")
		},
	}

	if v, ok := skills[*skillPara]; ok {
		v()
	} else {
		fmt.Println("not this skill")
	}

}
