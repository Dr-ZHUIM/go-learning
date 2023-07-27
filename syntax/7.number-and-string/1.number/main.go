package main

import (
	"fmt"
	"go-learning/utils"
)

func getAverage() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2

	fmt.Println(average)
}

func inaccuracyFloat() {
	ratio := 1.0 / 10.0
	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f\n", ratio)
}

func inaccuracyFloat2() {
	ratio := 3 / 2
	ratio2 := 3 / 2.0

	fmt.Printf("%d\n", ratio)
	fmt.Printf("%v\n", ratio2)
}

func main() {
	intWithFloat := 8 * -2.0
	intWithInt := 2 * 2
	utils.PrintType(intWithFloat, "intWithFloat")
	utils.PrintType(intWithInt, "intWithInt")
	getAverage()
	inaccuracyFloat()
	inaccuracyFloat2()
}
