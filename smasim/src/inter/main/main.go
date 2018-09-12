package main 

import (
	"io"
		"fmt"
		"inter"
)

func main() {
	//byteConter()
	interSatisf()
}




//IntSet is  testCode
type IntSet struct{ /*空的结构体*/}

func interSatisf() {
	//a := IntSet{}.String()	 //字面量无法取值;
	b := IntSet{}
	c := b.String() //变量是可以取值的
	fmt.Println(c)

	//数据处理.
	var w fmt.Stringer 
	w = &b  //Stringer 接口是 &IntSet实现的方法,而不是IntSet实现的方法;
	fmt.Printf("type :%T",w) //*IntSet

}


func (p *IntSet) String() string {
	return "test code"
} 


func byteConter() {
	//直接调用方法(注意在不同的包中,变量的引用要添加包名)
	var bc inter.ByteConter
    n, err := bc.Write([]byte{1,2,3,4,5}) 
    if err != nil {
		fmt.Printf("the error is %v \n",err)	
	}else {
		fmt.Printf("conter:%v\n",n)
	}

	//interface
	bc = 0
	fmt.Fprintf(&bc,"%s","hello")
	fmt.Printf("conter:%d\n",bc)

	var w io.Writer
	w = &bc  //Write是ByteConter指针的方法所以添加&
	fmt.Printf("type of bc : %T",w)

}