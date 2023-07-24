package main

import "fmt"

func main() {
	fmt.Printf("Hello!")
	test()
}

/*
	cannot redeclare the same function in the same package scope
*/
// func test() string {
// 	return "123"
// }
