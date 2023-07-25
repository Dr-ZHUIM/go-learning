package main

import (
	"fmt"
	"path"
)

func main() {
	var (
		num1  = 50
		num2  = 100
		dir   string
		file  string
		file2 string
	)
	num1, num2 = num2, num1
	num3, num4 := returnNums()
	dir, file = path.Split("assets/css/main.css")
	_, file2 = path.Split("assets/imgs/whatever.png")
	_, file3 := path.Split("assets/videos/whatever.mp4")
	fmt.Printf("dir: %v\n", dir)
	fmt.Printf("file: %v\n", file)
	fmt.Printf("file2: %v\n", file2)
	fmt.Printf("file3: %v\n", file3)
	println(num1, num2)
	println(num3, num4)
	num3, num4 = num4, num3
	println(num3, num4)

}

func returnNums() (int, int) {
	return 4, 3
}
