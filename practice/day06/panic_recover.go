package main

import "fmt"

func f1() { fmt.Println("f1") }
func f2() {
	defer func() { //必须把defer语句定义在panic语句之前
		err := recover() //recover之后 f3正常执行
		fmt.Printf("err:%v", err)
		fmt.Println("close fd")
	}()
	panic("fatal error")
	fmt.Println("f12")
}
func f3() { fmt.Println("f3") }
func main() {
	f1()
	f2()
	f3()
}
