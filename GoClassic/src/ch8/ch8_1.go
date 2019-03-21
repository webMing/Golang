package ch8
/// 课本ch7/http1
import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars)String() string{ 
	return fmt.Sprintf("$%.2f",d)
}

type dataBase map[string]dollars

func (db dataBase) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item,price := range db {
		fmt.Fprintf(w,"%s: %s\n",item,price)
	}
}

func Test() {
	//显示io.Buffer 到底是什么类型
	fmt.Println("This is Http1")
	// io.Writer 是接口类型 


}