package main

import (
	"fmt"
)

func main() {
	var apple int
	var orange int32
	apple = int(orange)
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)
	fmt.Println(string([]byte{104, 105}))
	_ = apple
}
