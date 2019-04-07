package ch10

import (
	"testing"
)

// 通过测试的标志
const passFlag = '\u2705'
// 没有通过测试的标志
const notPassFlag = '\u274E'

func Test_case1(t *testing.T) {
	t.Log("Testing content \n")
	t.Log(passFlag,notPassFlag)
	t.Logf("%c \n %c \n",passFlag,notPassFlag)
}

func Test_case2(t *testing.T){
	t.Log(" \n  Testing two \n")
}
