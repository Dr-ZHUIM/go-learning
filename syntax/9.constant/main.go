package main

import (
	"fmt"
	"go-learning/utils"
)

func magicNumber() {
	utils.PrintFuncName()
	const meter = 100
	cm := 100
	m := cm / meter
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meter
	fmt.Printf("%dcm is %dm\n", cm, m)
	utils.PrintBlank()
}

func useIota() {
	utils.PrintFuncName()
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	utils.PrintValue(monday, "monday")
	utils.PrintValue(tuesday, "tuesday")
	utils.PrintValue(wednesday, "wednesday")
	utils.PrintValue(thursday, "thursday")
	utils.PrintValue(friday, "friday")
	utils.PrintValue(saturday, "saturday")
	utils.PrintValue(sunday, "sunday")
	utils.PrintBlank()
}

func main() {
	magicNumber()
	useIota()
}
