package main

import "fmt"

func f1(f func()) {
	fmt.Println("f1")
	f()
}
func f2(a, b int) int {
	fmt.Println("f2")
	return a + b
}

func f(f2 func(int, int) int, a, b int) (f func()) {
	fmt.Println("func f")
	return func() { //note: 返回时没有执行该匿名函数，
		fmt.Println("anonymous func in f")
		f2(a, b)
	}
}
func main() {
	f1(f(f2, 2, 3))

}
