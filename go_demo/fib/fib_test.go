package fib

import (
	"fmt"
	"testing"
)

var c int

func TestFibList(t *testing.T) {
	// var a int =1
	// var a = 1
	// var b int =2

	// var (
	// 	a int = 1
	// 	b int = 1
	// )

	a := 1
	b := 1
	c = 1
	fmt.Println(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	a = 2
	a, b = b, a
	fmt.Println()
}
