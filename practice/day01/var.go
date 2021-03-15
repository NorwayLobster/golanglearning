package main

import "fmt"

var height int //可以不使用

func main() {
	//批量声明
	var (
		studentName string //小驼峰命名方式
		age         int
	)
	//非全局变量，必须使用
	var str string = "Go"
	var sex bool

	var sex1 bool = false
	var sex2 = false
	// 简短变量声明, 只能在函数中用
	sex3 := false
	//匿名变量： _
	// _, a=foo()
	fmt.Printf("name:%s", studentName)
	fmt.Println()
	fmt.Printf("age:%d\n", age)
	fmt.Println()
	fmt.Printf("height:%d", height)
	fmt.Println()

	fmt.Printf("sex:%v", sex)
	fmt.Println()

	fmt.Print(sex)
	fmt.Println()

	fmt.Printf("sex:%v", sex)
	fmt.Println()
	fmt.Printf("sex1:%v", sex1)
	fmt.Println()
	fmt.Printf("sex2:%v", sex2)
	fmt.Println()
	fmt.Printf("sex3:%v", sex3)
	fmt.Println()

	fmt.Printf("hello %s", str)

}
