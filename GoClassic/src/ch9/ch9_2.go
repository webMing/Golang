/**
 *	Go圣经 第八章 8.2习题
 */

package ch9

import (
	"os"
	"os/exec"
)

// Q2 测试
func Q2() {
	cmd := exec.Command("ls","-a")
	cmd.Path ="/tmp/"
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}