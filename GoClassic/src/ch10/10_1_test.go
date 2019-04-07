package ch10

import (
	"testing"
)

const passFlag = '\u2705'
const notPass = '\u274E'

func Test_case1(t *testing.T) {
	t.Log("Testing content \n")
	t.Log(passFlag,notPass)
	t.Logf("%c \n %c \n",passFlag,notPass)
}
