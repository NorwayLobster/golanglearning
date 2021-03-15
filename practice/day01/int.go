package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint = 10
	var b int = 10
	var c uintptr = 0x10000
	fmt.Printf("a:%T\n", a)
	fmt.Printf("b:%T\n", b)
	fmt.Printf("c:%T\n", c)

	var a1 int8
	var a2 int16
	var a3 int32
	var a4 int64

	var a5 uint8
	var a6 uint16
	var a7 uint32
	var a8 uint64

	fmt.Printf("a1:%T\n", a1)
	fmt.Println()
	fmt.Printf("a2:%T\n", a2)
	fmt.Println()
	fmt.Printf("a3:%T\n", a3)
	fmt.Println()
	fmt.Printf("a4:%T\n", a4)
	fmt.Println()
	fmt.Printf("a5:%T\n", a5)
	fmt.Println()
	fmt.Printf("a6:%T\n", a6)
	fmt.Println()
	fmt.Printf("a7:%T\n", a7)
	fmt.Println()
	fmt.Printf("a8:%T\n", a8)

	var i1 = 101 //int
	var i2 = 0x1234
	var i3 = 017
	var i4 = int8(17)

	fmt.Printf("%d\n", i1)
	fmt.Printf("%x\n", i1)
	fmt.Printf("%o\n", i1)

	fmt.Println()

	fmt.Printf("%d\n", i2)
	fmt.Printf("%x\n", i2)
	fmt.Printf("%o\n", i2)
	fmt.Println()

	fmt.Printf("%d\n", i3)
	fmt.Printf("%x\n", i3)
	fmt.Printf("%o\n", i3)

	fmt.Println()
	fmt.Printf("i4:%T\n", i4)

	fmt.Printf("MaxInt8:%d\n", math.MaxInt8)
	fmt.Printf("MinInt64:%d\n", math.MinInt64)
}
