package utils

import (
	"fmt"
)

func PrintType(variable any, key string) {
	fmt.Printf("typeof %s: %T \n", key, variable)
}
func PrintValue(variable any, key string) {
	fmt.Printf("valueof %s: %v \n", key, variable)
}
func PrintTypeAndValue(variable any, key string) {
	fmt.Printf("typeof %s: %T \nvalueof %[1]s: %v \n", key, variable, variable)
}
func PrintBlank() {
	fmt.Println("------------------------------")
}
