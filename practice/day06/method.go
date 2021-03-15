package main

import "fmt"

// type Dog struct { // 包外部可见
type dog struct {
	name string
	age  int
}

func (d dog) bark() { //接收者: 值拷贝,直接收者
	fmt.Println("wang wang wang")
	d.age++
}

func (ptr_d *dog) bark1() { //接收者: 指针接收者,
	// 使用场景：1. 结构体比较大，拷贝开销比较大时 2. 先回传值
	fmt.Println("wang wang wang")
	ptr_d.age++
}
func main() {
	//只能给自定义类型添加method
	//不能给基础类型、别的包中类型等添加方法,
	//但可以给基础类型的别名添加方法 e.g. type myInt int
	d := dog{"jonh", 2}
	d.bark()
	var ptr = &d
	fmt.Println(ptr)
	ptr.bark1()
	fmt.Println(ptr)

	var hi struct {
		h1 string
		h2 string
	}
	fmt.Println(hi)
	// hi.bark()
}
