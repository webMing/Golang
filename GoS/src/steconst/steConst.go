package steconst

import (
		"fmt"
)

//Some Values Same 有些值相同
const (
	SteAValue int = 1  //1
	SteBValue          //1 
	SteCValue          //1
	SteDValue  = "Name"  //Name
	SteEValue			//Name
)

//Diff Values 不同值
const (
	SteAAValue int = 1
	SteBBValue string = "Love"
	SteCCValue float64 = 34.4
)

//Emum Values 枚举值
const (
	SteAAAValue = 1 + iota
	SteBBBValue 
	SteCCCValue
)

// spci use
const (
	testValue uint8 = 0x01
	SteGValue       = testValue << iota
	SteHValue
	SteFValue
)

//"ChipType: custom defineType"
type  ChipType int32

//NONE 
const (
	NONE ChipType = iota
	CPU 
	GPU 
)

//trans to string
func (ct ChipType) String() (reV string) {
	switch ct {
	case NONE:{
			a := 1
			fmt.Println(a)
			reV = "none"
	}
	case GPU:
			reV = "gpu"
	case CPU:
			reV = "cpu"
	}
	return 
}

//ConstTest testCode
func ConstTest() {
	
	fmt.Printf("SteAValue:%d \n",SteAValue)
	fmt.Printf("SteCValue:%d \n",SteCValue)
	fmt.Printf("SteEValue:%s \n",SteEValue)
	fmt.Println("")

	fmt.Printf("SteAAValue:%+v \n",SteAAValue)
	fmt.Printf("SteBBValue:%+v \n",SteBBValue)
	fmt.Printf("SteCCValue:%+v \n",SteCCValue)
	fmt.Println("")

	fmt.Printf("SteAAAValue:%+v \n",SteAAAValue)
	fmt.Printf("SteBBBValue:%+v \n",SteBBBValue)
	fmt.Printf("SteCCCValue:%+v \n",SteCCCValue)
	fmt.Println("")

	fmt.Printf("SteGValue:%+v \n",SteGValue)
	fmt.Printf("SteHValue:%+v \n",SteHValue)
	fmt.Printf("SteFValue:%+v \n",SteFValue)
	fmt.Println("")
	
	fmt.Printf("SteHValue type : %T \n",SteHValue)

}