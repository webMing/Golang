package inird

import (
	"io"
		"bufio"
		"os"	
		"fmt"
		"strings"
)

//GetValue from  fileName at expectionSection which key is expect
func GetValue(fileName, expectSection, key string) (string,error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return fmt.Sprintf("GetValue:%v",err),err
	}
	value, err := readFile(file,expectSection,key)
	if err != nil {
		return "",err
	}
	return value,nil
}

//ReadFile readConTextFrom os.File
func readFile(f *os.File,sectionName,key string) (string,error) {
	reader := bufio.NewReader(f)
	var secName string
	for {

		lineStr, err := reader.ReadString('\n')

		if err == io.EOF {
			break //finish read!
		}

		if err != nil {
			return "",fmt.Errorf("read file failed: %v",err)
		}

		if lineStr == "" || lineStr == ";"{
			continue
		}
      
		lineStr = strings.TrimSpace(lineStr)

		if lineStr[0] == '[' && lineStr[len(lineStr)-1] == ']' {
			secName = lineStr[1:len(lineStr)-1]
		}else  if secName == sectionName {
			pair := strings.Split(lineStr,"=")
			if len(pair) == 2 {
				leftValue := strings.TrimSpace(pair[0])
				if leftValue == key {
					reValue := fmt.Sprintf("key:%v value:%v",key,strings.TrimSpace(pair[1]))
					return reValue,nil
				}
			}
		}
	}
	return "",nil

}