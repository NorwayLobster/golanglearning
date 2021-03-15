package main

import (
	"fmt"
)

//go中只有一种for语句,没有while和do whlie
func main() {
	var str1 string = "Hello 中国北京,"

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%v:%c\n", i, str1[i])
	}
	fmt.Printf("i:%d\t", i)

	fmt.Println()
	for i, c := range str1 {
		fmt.Printf("%v:%c\n", i, c)
	}

	fmt.Println()
	var i = 0
	for ; i < 10; i++ {
		if i > 8 {
			break
		}
		fmt.Printf("i:%d\t", i)
	}
	fmt.Println()
	fmt.Printf("i:%d\t", i)

	fmt.Println()
	var ii = 0
	for ii < 10 {
		ii++
		if i == 5 {
			continue
		}
		fmt.Printf("ii:%d\t", ii)
	}
	fmt.Println()

	// for { fmt.Printf("ii:%d", ii) } //dead circle
}
