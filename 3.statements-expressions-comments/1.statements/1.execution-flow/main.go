package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")
	num := 0
	for num < 10 {
		fmt.Printf("num: %v\n", num)
		if num == 9 {
			fmt.Printf("num: %v", num)
		}
		num++
	}
}
