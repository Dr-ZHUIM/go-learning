package main

import (
	"fmt"
	"go-learning/utils"
	"math"
	"time"
	"unsafe"
)

type (
	gram  float64
	ounce float64
)

func printBits() {
	fmt.Printf("%02b = %d\n", 0, 0)
	fmt.Printf("%02b = %d\n", 1, 1)
	fmt.Printf("%02b = %d\n", 2, 2)
	fmt.Printf("%02b = %d\n", 3, 3)
	utils.PrintBlank()
}

func printbytes() {
	fmt.Printf("%08b = %d\n", 1, 1)
	fmt.Printf("%08b = %d\n", 2, 2)
	fmt.Printf("%08b = %d\n", 4, 4)
	fmt.Printf("%08b = %d\n", 8, 8)
	fmt.Printf("%08b = %d\n", 16, 16)
	fmt.Printf("%08b = %d\n", 32, 32)
	fmt.Printf("%08b = %d\n", 64, 64)
	fmt.Printf("%08b = %d\n", 128, 128)
	utils.PrintBlank()
}

func predeclaredTypes() {
	fmt.Println("int8   :", math.MinInt8, math.MaxInt8)
	fmt.Println("int16  :", math.MinInt16, math.MaxInt16)
	fmt.Println("int32  :", math.MinInt32, math.MaxInt32)
	fmt.Println("int64  :", math.MinInt64, math.MaxInt64)

	// unsigned type can only represent positive integers
	fmt.Println("uint8  :", 0, math.MaxUint8)
	fmt.Println("uint16 :", 0, math.MaxUint16)
	fmt.Println("uint32 :", 0, math.MaxUint32)
	fmt.Println("uint64 :", 0, uint64(math.MaxUint64))
	// You can't print MaxUint64 directly
	// Its size is way bigger for the runtime
	// It can only be used in constant expressions

	// e means zeros after the number (scientific notation)
	// 1e1, is 10 1e2, is 100, 12e1 is 120
	fmt.Println("float32:", math.SmallestNonzeroFloat32,
		math.MaxFloat32)
	fmt.Println("float64:", math.SmallestNonzeroFloat64,
		math.MaxFloat64)

	// memory costs
	fmt.Println("int    :", unsafe.Sizeof(int(1)), "bytes")
	fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "bytes")
	fmt.Println("int16  :", unsafe.Sizeof(int16(1)), "bytes")
	fmt.Println("int32  :", unsafe.Sizeof(int32(1)), "bytes")
	fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "bytes")

	fmt.Println("uint   :", unsafe.Sizeof(uint(1)), "bytes")
	fmt.Println("uint8  :", unsafe.Sizeof(uint8(1)), "bytes")
	fmt.Println("uint16 :", unsafe.Sizeof(uint16(1)), "bytes")
	fmt.Println("uint32 :", unsafe.Sizeof(uint32(1)), "bytes")
	fmt.Println("uint64 :", unsafe.Sizeof(uint64(1)), "bytes")

	fmt.Println("float32:", unsafe.Sizeof(float32(1)), "bytes")
	fmt.Println("float64:", unsafe.Sizeof(float64(1)), "bytes")

	fmt.Println("hello  :", len("hello")+8, "bytes")
	fmt.Println("hi     :", len("hi")+8, "bytes")
	utils.PrintBlank()

}

func convertTime() {
	date, _ := time.ParseDuration("4h30m")
	fmt.Println(date.Hours(), "hours")
	fmt.Println(date.Minutes(), "minutes")

	var m int64 = 2
	date *= time.Duration(m)
	fmt.Println(date)
}

func useType() {
	var g gram = 1000                 // gram  -> float64
	var o ounce = ounce(g) * 0.035274 // ounce -> float64

	fmt.Printf("%g grams is %.2f ounce\n", g, o)
}

func main() {
	printBits()
	printbytes()
	predeclaredTypes()
	convertTime()
	useType()
}
