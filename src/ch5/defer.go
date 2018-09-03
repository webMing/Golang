package main

import (
	"fmt"
	"os"
	)

func main() {
	fmt.Println("---------------------")
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
	fmt.Println("---------------------")
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f,"custom Data")
}
