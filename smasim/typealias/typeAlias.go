package typealias
/*
	类型定义 和 类型别名一定要区别

	类型定义形式(重新定义一个类型和原来的类型不同):
	type SteInt int

	类型别名(和原来的类型是相同，只是有了一个小名):
	type Ste_int = int 

*/
import (
	"fmt"
)

//类型别名
type IntAlias = int

//类型定义
type NewInt int

//SteTestTypeAlias
func SteTestTypeAlias(){
	fmt.Println("test begin")
	var a IntAlias 
	var b NewInt	
	fmt.Printf("%T \t %T \n",a,b)
	fmt.Println("test end")
}
