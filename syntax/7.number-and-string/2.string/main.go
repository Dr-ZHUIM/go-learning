package main

import (
	"fmt"
	"go-learning/utils"
	"strconv"
	"strings"
	"unicode/utf8"
)

func assignString() {
	var s string
	s = "I'm a string"
	s = `I'm a string`
	utils.PrintTypeAndValue(s, "s")
	utils.PrintBlank()
}

func converNonString() {
	boola, boolb := true, false
	num1, num2 := 1, 2
	str1 := "1 + 2 = "
	str2 := "bool has "
	utils.PrintValue(str1+strconv.Itoa(num1+num2), "convertNumberToString")
	utils.PrintValue(str2+strconv.FormatBool(boola)+" and "+strconv.FormatBool(boolb), "convertBoolToString")
	utils.PrintBlank()
}

func stringLength() {
	str := "I'm a string"
	utils.PrintValue(len(str), "length of "+str)
	utils.PrintBlank()
}

func unicodeLength() {
	name := "İnanç"
	name2 := "万维网"
	utils.PrintValue(len(name), "bytes length of "+name)
	utils.PrintValue(utf8.RuneCountInString(name), "characters length of "+name)
	utils.PrintValue(len(name2), "bytes length of "+name2)
	utils.PrintValue(utf8.RuneCountInString(name2), "characters length of "+name2)
	utils.PrintBlank()
}

func windowPath() {
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)
	utils.PrintBlank()
}

func printJson() {
	json := `
{ 
	"Items": [{ 
		"Item": {
			"name": "Teddy Bear"
		} 
	}] 
}`
	fmt.Println(json)
	utils.PrintBlank()
}

func lower() {
	str := "SHEPARD"
	fmt.Println(strings.ToLower(str))
	utils.PrintBlank()
}

func trim() {
	msg := `
	
	The weather looks good.
I should go and play.



	`
	fmt.Println(strings.TrimSpace(msg))
	utils.PrintBlank()
}

func trimRight() {
	name := "inanç           "
	name = strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(name)
	fmt.Println(l)
	utils.PrintBlank()
}

func main() {
	assignString()
	converNonString()
	stringLength()
	unicodeLength()
	windowPath()
	printJson()
	lower()
	trim()
	trimRight()
}
