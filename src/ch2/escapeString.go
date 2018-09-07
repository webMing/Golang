package main

import ("fmt")
	

func main() {
        //转义符转义	
	escape := "c:\\go\\bin\\go"
	fmt.Println(escape)
	
	//元字符串
	rawStr := `first 
		  \n \n \n
		   second
		  `
	fmt.Println(rawStr)	
}
