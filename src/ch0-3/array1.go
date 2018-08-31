package main

import ("fmt")

func main() {
	fmt.Printf("The First Code!")
	var a [3]int
	fmt.Printf("%v \n",a)

	fmt.Println("a[2] is :",a[2])	
	fmt.Println("len(a) is :",len(a))
	
	for i := 0 ; i < len(a) ;i++ {
		fmt.Printf("a[%d] = %d \n",i,a[i])
	}
	for index,value := range a {
		fmt.Printf("a[%d] = %d \n",index,value)
	}

	var b [4]string
	b[0] = "a"
	b[2] = "b"	
	for index,value := range b {
		fmt.Printf("b[%d] = %s \n",index,value)
	}
	
	//var c [4]rune
	//var c = [3]rune{'a','b','c'}
	var c = [...]rune{'c','d','e','f','e'}
	c[2] = 'C'
	for index,value := range c {
		fmt.Printf("c[%d] = %q \n",index,value)
	}

	d := [...]float64{One:1.0,Two:2.0,Three:3.0}
	for index,value := range d {
		fmt.Printf("d[%d] = %f \n",index,value)
	}
	
	e := [...]string{99:"99"}
	for index,value := range e {
		fmt.Printf("e[%d] = %s \n",index,value)
	}

	//数组比较
	f := [...]int{1,2}
	g := [2]int{1,2}
	var k = [...]int{1,3}
	fmt.Printf("f == g %v, f == k %v , g == k %v \n",f==g,f == k,g ==k)


	//数据创建
	var h [5]string
	h[1] = "h[1]"
	h[3] = "h[3]"
	for index,value := range h {
		fmt.Printf("h[%d] = %s \n",index,value)
	}	
	
	change(&h)
	for index,value := range h {
		fmt.Printf("h[%d] = %s \n",index,value)
	}

	//h := [3]int
	//h := [...]int{1,3}
	//h := [...]int{10:4}
}


type Current int

const (
	One Current = 1 + iota
	Two
	Three
      )

//传递指针完成值的改变
func change(a *[5]string) {
	a[1] ="dfd"
	a[2] ="dfd" 
}
