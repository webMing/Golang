package inter

import (
	"time"
	"fmt"
	"flag"
)

//返回值是指向time.Duration的指针
var period = flag.Duration("period", 1 * time.Second,"sleep period")
//FlagValue is test func
func FlagValue() {
	flag.Parse() //解析
	fmt.Printf("Sleeping for %v...\n",*period)
	time.Sleep(*period)
	fmt.Println()
}

