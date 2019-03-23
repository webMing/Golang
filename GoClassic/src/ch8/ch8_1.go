package ch8
/// 课本ch7/http1
import (
	"fmt"
	"net/http"
)

// Ste test type
type Ste int 
// Aab 是一个测试方法
func (s Ste) ab(){
// 这个是空的方法，
}

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


