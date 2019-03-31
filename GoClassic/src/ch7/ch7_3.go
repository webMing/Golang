package ch7
import (
	"time"
	"fmt"
	"flag"
)
var period = flag.Duration("period", 1 * time.Second, "sleep peroid")	
//Q1 解析命令行参数
func Q1() {
	fmt.Println(period)
	flag.Parse()
	fmt.Printf("Sleeping for %v....",*period)
	time.Sleep(*period)
	fmt.Println()
}