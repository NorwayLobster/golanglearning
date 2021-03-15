package main

import (
	"fmt"
	"math"
)

func main() {
	var a0 = 12.2 //默认float64
	var a1 float32 = math.MaxFloat32
	var a2 float32 = math.SmallestNonzeroFloat64

	var a3 = float32(12.2) //默认float64, 强转
	a1 = a2
	// a1 = a0 //error
	fmt.Printf("a0:%T\n", a0)
	fmt.Printf("a0:%f\n", a0)
	fmt.Println()

	fmt.Printf("a1:%T\n", a1)
	fmt.Printf("a1:%f\n", a1)
	fmt.Println()
	fmt.Printf("a2:%T\n", a2)
	fmt.Printf("a2:%f\n", a2)

	fmt.Println()
	fmt.Printf("a3:%T\n", a3)
	fmt.Printf("a3:%f\n", a3)
}
