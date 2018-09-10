package ctlflow

import "fmt"

//forBreakExample 利用条件变量判断跳出for循环
func forBreakExample() {
	 breakOut := false 
	 for i := 0 ;;i++ {
		 for j := 0; ;j++ {
			 if j > 10 && i > 3 {
				 breakOut = true
				 break 
			 }
		 }
		 if breakOut {
			 break 
		 }
	 }
}

//可以使用goto跳出多重循环

//GotoStatement 使用goto跳出重循环
func GotoStatement() {
	for i := 0 ; ;i++ {
		for j := 0 ; ; j++ {
			if j > 10 && i > 3 {
				goto label
			}
		}
	}
	//Label标签
	label: customPrintf()
}

func customPrintf() {
	fmt.Println("customPrintf")
}

