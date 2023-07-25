package main

import (
	"fmt"
	"os"
)

func person(index int) string {
	suffix := " name -"
	switch index {
	case 1:
		return "1st" + suffix
	case 2:
		return "2nd" + suffix
	case 3:
		return "3rd" + suffix
	default:
		return fmt.Sprint(index) + "ed" + suffix
	}
}

func main() {
	// fmt.Printf("%#f\n", os.Args[0])
	count := len(os.Args) - 1
	fmt.Println("path -", os.Args[0])
	fmt.Printf("There are %v names.\n", count)
	fmt.Println(person(1), os.Args[1])
	fmt.Println(person(2), os.Args[2])
	fmt.Println(person(3), os.Args[3])
	fmt.Println(person(4), os.Args[4])
	fmt.Println(person(5), os.Args[5])
}
