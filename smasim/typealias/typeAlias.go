package typealias

import (
	"fmt"
	"time"
)
/*
	类型定义 和 类型别名一定要区别

	类型定义形式(重新定义一个类型和原来的类型不同):
	type SteInt int

	类型别名(和原来的类型是相同，只是有了一个小名):
	type Ste_int = int 

*/

//类型别名
type IntAlias = int

//类型定义
type NewInt int

//SteTestTypeAlias
func SteTestTypeAlias(){
	//区别不同之处
	var a IntAlias 
	var b NewInt	
	fmt.Printf("%T \t %T \n",a,b)
}

//如果我想给 myDuration 添加方法是不行的,因为time.Duration和typeAlias相比属于不同的包
type myDuration = time.Duration
/*下面这样会报错*/
//func (m myDuration) doSome() {}