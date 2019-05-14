package storage 

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

//Stor1 is....
func Stor1() {
	
	// 判断当前文件是否已经存在
	fileInfo,err := os.Stat("data1")
	fileExist := true	
	if err == nil {
		fmt.Printf("File Name:%s", fileInfo.Name())
	}else if os.IsNotExist(err) {
		fmt.Println("data1 file not exist")
	    fileExist = false
	}else {
		fileExist = false
		log.Fatal(err) //other reson
	}

	if !fileExist {
		data := []byte("File is very good Hello World! \n")
		// 如果文件存在就会覆盖之前文件中的内容
		err = ioutil.WriteFile("data1",data,0644)
		checkErr(err)
	} 

	//从data1 文件中读取数据
	read1,err := ioutil.ReadFile("data1")
	checkErr(err)
	fmt.Print(string(read1))

}

// checkErr checkErr用来检查数据
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}