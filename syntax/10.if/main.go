package main

import (
	"go-learning/utils"
)

func boolOperator() {
	utils.PrintFuncName()
	speed := 100
	isFast := speed >= 80
	isSlow := speed <= 80
	utils.PrintValue(isFast, "isFast")
	utils.PrintValue(isSlow, "isSlow")
	utils.PrintBlank()

}

func main() {
	boolOperator()
}
