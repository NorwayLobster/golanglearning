package main

import "fmt"

func main() {
	// new: 一般用于给基本数据类型分配内存
	var a int = 10
	var p *int = &a
	fmt.Println(a, p)
	*p = 11
	fmt.Println(a, p)
	var p1 *int = new(int)
	fmt.Println(*p1, p1) //对应地址的零值
	*p1 = 1000
	fmt.Println(*p1, p1)

}
