package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

//iota:
func main() {
	//case1:
	const (
		a1 = iota //实现类似枚举
		_
		a2
		a3
	)
	fmt.Printf("a1:%d\n", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println()
	//case2:
	const (
		b1 = iota //实现类试枚举
		b2 = iota
		b3
	)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println()
	//case3:
	const (
		c1 = iota //实现类试枚举
		c2 = 100
		c3 = iota
		c4
	)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println()
	//case4:
	const (
		d1, d2 = iota + 1, iota + 2
		d3, d4 = iota + 1, iota + 2
	)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

}
