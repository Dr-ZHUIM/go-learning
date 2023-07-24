package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	const ok = true
	var hello = "Hello"
	var (
		name1 = "TOM"
		name2 = "BOM"
	)
	name3, name4 := swap(name1, name2)
	num1, num2 := split(17)
	fmt.Println(hello, name3, name4)
	fmt.Println(hello, name1, name2)
	fmt.Printf("num1: %v\n", num1)
	fmt.Printf("num2: %v\n", num2)
}
