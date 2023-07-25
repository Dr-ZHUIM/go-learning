package main

import (
	"fmt"
)

func main() {

	a := 1
	b := 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
