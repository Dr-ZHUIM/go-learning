package main

import (
	"fmt"
	"go-learning/syntax/4.write-your-first-library-package/zhuim"
)

func main() {
	zhuim.PrintZhuim()
	fmt.Printf("zhuim.IsMobile: %v\n", zhuim.IsMobile)
	version := zhuim.CallVersion()
	fmt.Printf("version: %v\n", version)
}
