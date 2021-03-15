package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	var a2 []int
	fmt.Printf("a1:%v,len:%v,cap:%v\n", a1, len(a1), cap(a1))
	fmt.Printf("a2:%v,len:%v,cap:%v\n", a2, len(a2), cap(a2))
	copy(a2, a1)
	fmt.Printf("a1:%v,len:%v,cap:%v\n", a1, len(a1), cap(a1))
	fmt.Printf("a2:%v,len:%v,cap:%v\n", a2, len(a2), cap(a2))
	a3 := make([]int, 4, 20)

	fmt.Printf("a3:%v,len:%v,cap:%v\n", a3, len(a3), cap(a3))
	copy(a3, a1)
	fmt.Printf("a1:%v,len:%v,cap:%v\n", a1, len(a1), cap(a1))
	fmt.Printf("a3:%v,len:%v,cap:%v\n", a3, len(a3), cap(a3))
}
