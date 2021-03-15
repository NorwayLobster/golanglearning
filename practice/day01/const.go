package main

import "fmt"

func main() {
	const pi = 3.14
	fmt.Println("pi:", pi)

	const (
		statusOK  = 200
		notFound  = 404
		notFound1 //批量声明时，如果某一行声明后没有初始化，默认和上一行一样的值
	)

	fmt.Println("statusOK:", statusOK)
	fmt.Println("notFound:", notFound)
	fmt.Println("notFound1:", notFound1)

}
