package main

import "fmt"

func main() {
	//make: 用于给slice，map，chan申请内存，但不返回指针，而是返回该对象本身
	//注意与new的区别
	var len1 int = 3
	var cap1 int = 10
	a1 := make([]int, len1, cap1)
	fmt.Printf("a1:%v, len:%v, cap:%v\n", a1, len(a1), cap(a1))

}
