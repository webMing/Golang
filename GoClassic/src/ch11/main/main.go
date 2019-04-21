package main

import (
	"fmt"
	"ch11"
)
type myTT struct{}

func (t *myTT) y(c int) {
	//do nothing
}

type kf func(c int)

func main() {
	s := "/a/b/c/d!"
	sSlice := s[1:]
	fmt.Println(sSlice)
	
	t := myTT{}
    handle(new(myTT))
	// ch11.Extern1()	
	// ch11.Extern2()
	// ch11.Extern3()
	ch11.Extern4()	
}

func handle(k kf) {

}