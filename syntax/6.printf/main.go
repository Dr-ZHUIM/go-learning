package main

import (
	"fmt"
	"strings"
)

func getType(variable any, key string) {
	fmt.Printf("typeof %v: %T\n", key, variable)
}

func getValue(variable any, key string, unit string) {
	fmt.Printf("%v: %v %v\n", strings.ToTitle(key), variable, unit)
}

func printType() {
	var (
		brand string
		speed int
		heat  float64
		off   bool
	)
	fmt.Println("----------------printType----------------")
	getType(brand, "brand")
	getType(heat, "heat")
	getType(speed, "speed")
	getType(off, "off")
}

func coding() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Println("----------------coding----------------")
	getValue(planet, "planet", "")
	getValue(distance, "distance", "millions kms")
	getValue(orbital, "orbital", "days")
	getValue(hasLife, "hasLife", "")
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}

func main() {
	ops, ok, fail := 2350, 543, 433
	var (
		brand string
	)
	fmt.Println("total:", ops, "- success:", ok, "/", fail)
	fmt.Printf("total: %v - success: %v / %v \n", ops, ok, fail)
	fmt.Printf("brand: %q\n", brand)
	brand = "google"
	fmt.Printf("brand: %q\n", brand)
	fmt.Println("hi\"hi\"")
	printType()
	coding()
}
