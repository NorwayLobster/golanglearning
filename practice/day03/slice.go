package main

import "fmt"

//包级别的变量初始化过程
var a = b + c
var b = 2
var c = 1

func main() {
	fmt.Printf("a:%d\n", a)
	fmt.Printf("b:%d\n", b)
	fmt.Printf("c:%d\n", c)
	//define  a slice
	var a0 []int           //切片的定义，底层实现就是一个数组
	fmt.Println(a0 == nil) //true

	fmt.Println()
	var s1 = []int{}
	fmt.Println(s1 == nil) //false
	fmt.Printf("s1's len:%d\n", len(s1))
	fmt.Printf("s1's cap:%d\n", cap(s1))

	fmt.Println()
	fmt.Printf("a0's len:%d\n", len(a0))
	fmt.Printf("a0's cap:%d\n", cap(a0))
	fmt.Printf("a0:%v\n", a0)
	a0 = []int{1, 2, 4}
	fmt.Println("after init:")
	fmt.Printf("a0's len:%d\n", len(a0))
	fmt.Printf("a0's cap:%d\n", cap(a0))
	fmt.Printf("a0:%v\n", a0)

	fmt.Println()

	fmt.Println(s1 == nil)

	fmt.Println()
	//init a slice
	var a3 []int //切片的定义，底层实现就是一个数组
	fmt.Println(a3)
	a31 := [...]int{1, 2, 3, 4, 5, 6, 7, 8} //切片的定义，底层实现就是一个数组
	a3 = a31[2:3]
	fmt.Println(a3)
	fmt.Printf("a3's len:%d\n", len(a3))
	fmt.Printf("a3's cap:%d\n", cap(a3))
	//
	fmt.Println()

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8} //切片的定义，底层实现就是一个数组
	a4 := a1[2:5]
	fmt.Printf("a4's cap:%d\n", cap(a4))
	fmt.Printf("a4's len:%d\n", len(a4))
	fmt.Printf("a4:%v\n", a4)

	fmt.Println()

	//由切片得到切片
	a5 := a4[2:3]                        //此index 2是slice a4的index，不是底层数组a1的index
	fmt.Printf("a5's cap:%d\n", cap(a5)) //
	fmt.Printf("a5's len:%d\n", len(a5))
	fmt.Printf("a5:%v\n", a5)

	a6 := a4[2:]
	fmt.Printf("a6's cap:%d\n", cap(a6)) //
	fmt.Printf("a6's len:%d\n", len(a6))
	fmt.Printf("a6:%v\n", a6)

	fmt.Println()
	//可以通过切片的修改底层数组
	//切片的本质：底层数组的某一个视图的引用，底层实现是一个聚合数据结构：ptr，capacity,length
	fmt.Printf("a6[0]:%v\n", a6[0])
	fmt.Printf("a1:%v\n", a1)
	a6[0] = 10000
	fmt.Printf("a6[0]:%v\n", a6[0])
	fmt.Printf("a1:%v\n", a1)

	//slice capacity: 切片指向一个底层数组，切片的容量: 切片开始位置到底层数组的最后元素
}
