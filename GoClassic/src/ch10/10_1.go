package ch10
/*
	对应到GO In Action 中, 关于log包的内容
*/
import (
	"log"
)

func simpleLog() {
	log.Print("Simple log \n")	
	log.Printf("Simple log %s \n","S")	
	log.Println("Simple log","you kown")
}