package main

import "fmt"

func main() {
	var s1 []int
	//动态扩容策略:
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))
	s2 := []int{5, 6, 7}
	s1 = append(s1, s2...) //...表示拆包，展开
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))

	//匿名切片+拆包
	s1 = append(s1, []int{1, 3}...) //...表示拆包，展开
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))

	//切片的变相删除，即覆盖
	s1 = append(s1[:1], s1[2:]...) //把index为1的元素覆盖
	fmt.Println("after delete:")
	fmt.Printf("s1:%v, len:%v,cap:%v\n", s1, len(s1), cap(s1))
	fmt.Println(s1, len(s1), cap(s1))

	//case1:
	fmt.Println()
	var a1 = make([]int, 5, 10)
	fmt.Println(a1, len(a1), cap(a1))
	for i := 0; i < 10; i++ {
		a1 = append(a1, i)
	}
	fmt.Println(a1, len(a1), cap(a1))

	//case2:
	fmt.Println()
	var a2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("a2:%v, len:%v,cap:%v\n", a2, len(a2), cap(a2))
	fmt.Printf("s2:%v, len:%v,cap:%v\n", s2, len(s2), cap(s2))
	s2 = append(a2[:1], a2[2:]...) //把index为1的元素覆盖
	fmt.Println("after delete:")
	// fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("a2:%v, len:%v,cap:%v\n", a2, len(a2), cap(a2))
	fmt.Printf("s2:%v, len:%v,cap:%v\n", s2, len(s2), cap(s2))
}
