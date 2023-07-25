package zhuim

import (
	"fmt"
	"runtime"
)

var (
	IsMobile = false
	IsWebApp = true
)

func PrintZhuim() {
	fmt.Println("Zhuim")
}

func CallVersion() string {
	return runtime.Version()
}

func init() {
	fmt.Println("package zhuim has initialized!")
}
