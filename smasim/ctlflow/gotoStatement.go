package ctlflow

import "fmt"

//forBreakExample 利用条件变量判断跳出for循环
func forBreakExample() {
	 breakOut := false 
	 for i := 0 ;;i++ {
		 for j := 0;j < 15 ;j++ {
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

/*
可以使用goto跳出多重循环
利用goto语句也可以对错误进行统一的处理
*/

//GotoStatement 使用goto跳出重循环
func GotoStatement() {
	for i := 0 ; ;i++ {
		for j := 0 ;j < 7 ; j++ {
			if (j > 2) && (i > 2) {
				goto label
			}
			fmt.Printf("goto test i:%d j:%d \n",i,j)
		}
	}
	//Label标签
	label: customPrintf()
}

func customPrintf() {
	fmt.Println("customPrintf")
}

func continueExcute() {
	
	for i := 0 ; i < 20 ;i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf("i : %d \n",i)
	}	
}
