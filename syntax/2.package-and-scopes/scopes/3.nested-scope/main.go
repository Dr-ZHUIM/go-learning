package main

import "fmt"

var declareMeAgain = 20

func nested() {
	declareMeAgain := 12
	fmt.Println("inside nested:", declareMeAgain)
}

func main() {
	fmt.Println("inside main:", declareMeAgain)
	nested()
	fmt.Println("inside main:", declareMeAgain)
}
