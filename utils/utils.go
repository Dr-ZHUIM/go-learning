package utils

import (
	"fmt"
	"runtime"
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
func PrintFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Printf("func %s:\n\n", f.Name())
}
