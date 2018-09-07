package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
	)

func main() {

	fmt.Println("========================")
	
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int 
	invalid := 0
		
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr,"charCount: %v \n",err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	
	fmt.Printf("rune \t count \n")
	for c,n := range counts {
		fmt.Printf("%q\t%d\n",c,n)
	}
	fmt.Printf("\nlen \t count \n")
	for l,c := range utflen {
		if l > 0 {
             		fmt.Printf("%d\t%d\n",l,c)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n %d invalid UTF-8 characters\n",invalid)
	}		
	fmt.Println("========================")
}
