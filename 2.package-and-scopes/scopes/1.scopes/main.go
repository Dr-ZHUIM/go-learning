package main

// file scope begins
/*
file scope can only work in this file
*/
import "fmt"

// file scope ends

// package scope begins

const str = "some literals"

func main() {
	// block scope begins
	const blockStr = "this string can only be called in main function!"
	fmt.Printf("blockStr: %v\n", blockStr)
	fmt.Printf("str: %v\n", str)
	// block scope ends
}
