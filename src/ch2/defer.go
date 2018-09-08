package main

import (
	"fmt"
	"os"
	//"unicode/utf8"
	"time"
)

func main() {
	fmt.Println("Hello world!") //方法调用
	//deferTest1()
	//captureVar1()
	//captureVar2()
	//deferTest3()
	otherCode3()
 }

// defer 简单测试
func deferTest1() {
	f := fileCreate("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func fileCreate(p string) (f *os.File) {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("wrting")
	fmt.Sprintln(f, "data")
}

//defer 测试2
func deferTest2() {
	//
	s := []int{8, 9, 10,11,12}
	for i := 0; i < len(s); i++ {
		/*
		*  defer 后面跟的是function；延时执行的也是这个function
		*  defer 执行是在：1.函数执行完所有的定义指令 2. 如果有 panic 实在panic执行完之后 3. return 语句之后
		*  defer 延迟执行函数
		*  defer 是在执行defer 后的函数的时候进行参数以及返回值的评估
		* defer  后面function值的评估是在 defer语句进行的时候执行; 其中评估的值包括参数和函数返回值
		* defer 后要求是函数的执行如果返回的是匿名函数则需要添加（）
		*/
	 	defer fmt.Printf("i value is : %d \n",i)
	}
	fmt.Println("func reach end!") //defer 是在这条语句执行之后再执行
}

//defer 测试2
func deferTest3() {
	//
   defer track()()
   fmt.Printf("func in deferTest3! \n")
}

func track() func() {
	fmt.Printf(" I am excuting ! \n")
	return func(){
		fmt.Printf("deferTest3 reach end! \n")
	}
}

// 测试匿名函数捕获功能1
func captureVar1() {
	var t = 3
	fmt.Printf("t value:%d \n", t)
	f := func() { t = 5 }
	f()
	fmt.Printf("t value:%d \n", t)
	t = 23
	fmt.Printf("t value:%d \n", t)

}

// 测试匿名函数捕获2
func captureVar2() {
	s := []func(){}
	for _, v := range []int{1, 3, 4, 5, 6, 7} {
		//由于匿名函数捕获的是变量的地址所以才会显示都是7;
		v := v
		s = append(s, func() { fmt.Printf("i value :%d \n", v) })
	}
	for i := 0; i < len(s); i++ {
		s[i]()
	}
}


func otherCode1() {

	/*
	for _,v :=  range []int{23,45,56,67,78} {
	   //匿名函数会捕捉到v这个变量地址,而不是v这个值, 注意这里是v变量地址的捕捉,而没有对值进行评估,defer会对参数，以及返回值进行评估	
		defer func(){
			fmt.Printf("v value: %d \n",v)
		}()
	}
	*/

	/*
	for _,v :=  range []int{23,45,56,67,78} {
		//当执行defer func 语句的时候会对参数,以及返回值进行评估.
		defer func(n int){
			fmt.Printf("n value: %d \n",n)
		}(v)
	}
	*/

	/*证明参数以及返回值值的评估是在执行defer语句的时候进行
	var x = 3
	defer fmt.Printf("the value of x is:%d \n",x)
	x = 4 
	fmt.Printf("the value of x is : %d \n",x)	
    */
	fmt.Println("reach func end!")

	//https://stackoverflow.com/questions/24720097/golang-defer-behavior
}

//书中代码测试
func otherCode3() {
	defer otherCodeTrack3()() //注意这里的（）别丢了
	time.Sleep(3 * time.Second)
}

func otherCodeTrack3() func() {
	t := time.Now()
	fmt.Printf("other code3 begin excute! \n")
	return func(){
		fmt.Printf("otherCode3 excute time %s \n",time.Since(t).String())
	}
}

