package tempconv

import (
	"flag"
)

//CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

//FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9 )
}

//CelsiusFlag defines a Celsius flag with the specified name.
func CelsiusFlag(name string,value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f,name,usage)
	return &f.Celsius
}