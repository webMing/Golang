package tempconv

import	"fmt"

//Celsius temperture
type Celsius  float64 //摄氏度
//Fahrenheit temperture
type Fahrenheit float64 //华氏度
//CelsiusFlag satisfies the flag.Value interface
type celsiusFlag struct {Celsius}

//AboluteZeroC FreezingC BoilingC
const (
	 AbsoluteZeroC  Celsius = -273.15  //绝对零度
	 FreezingC 		Celsius = 0      //凝点
	 BoilingC       Celsius = 100   //沸点
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C",c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F",f)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s,"%f%s",&value,&unit) //为什么是地址尼
	switch unit {
		case "C","°C" : 
			f.Celsius = Celsius(value)
			return nil
		case "F","°F" :
			f.Celsius = FToC(Fahrenheit(value))
			return nil
	}
	return fmt.Errorf("invalid temperature %q",s)
}
