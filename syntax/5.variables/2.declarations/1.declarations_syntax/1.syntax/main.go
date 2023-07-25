package main

import "fmt"

var name string

func main() {
	var speed int
	speed = 23
	// redclaration only works when one variable is new
	speed, color := 50, "red"
	fmt.Printf("speed: %v\n", speed)
	fmt.Printf("color: %v\n", color)
}
