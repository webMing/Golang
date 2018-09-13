package main 

import (
	"bytes"
	"io"
	"fmt"
)

const  debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	// do other task

}

//接口转换出现错误
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done ! \n"))
	}
}

//emptyInter 空接口类型
func emptyInter(){
   var a interface{}
   a = nil
   a = 2
   a = struct{}{}
   a = 34.3
   a = [4]int{0}
   a = []int{34:0}
   a = map[string]int{}
	fmt.Println(a)

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

