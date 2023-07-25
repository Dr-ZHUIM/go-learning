package main

import (
	"fmt"
)

func main() {
	speed := 100
	force := 2.5
	fmt.Printf("speed: %T\n", speed)
	fmt.Printf("speed: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)
	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
