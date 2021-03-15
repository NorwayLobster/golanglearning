package main

import "fmt"

func main() {
	//init:

	//case1:
	var arr1 [3]bool
	arr1 = [3]bool{true, true, true}
	fmt.Print(arr1)
	fmt.Println()

	//case2:
	var arr2 [4]bool
	arr2 = [...]bool{true, false, true, true}
	fmt.Print(arr2)
	fmt.Println()

	//case3:
	var arr3 [4]int
	arr3 = [4]int{1, 2, 3, 4}
	fmt.Print(arr3)
	fmt.Println()

	//case4:
	var arr4 [5]int
	arr4 = [5]int{0: 1, 3: 4} //根据索引指定init
	fmt.Print(arr4)
	fmt.Println()

	b1 := [3]int{1, 2, 3}
	b2 := b1 //值类型
	b2[1] = 100
	fmt.Println(b1)
	fmt.Println(b2)

	var c1 [2][3]int
	c1 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(c1)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", c1[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	for _, v := range c1 {
		// fmt.Print(v)
		for _, e := range v {
			fmt.Printf("%d\t", e)
		}
		fmt.Println()
	}
}
