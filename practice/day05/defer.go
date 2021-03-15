package main

import "fmt"

func f(a int) {
	fmt.Printf("%v\n", a)
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //副本++
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}() //副本++
	return x
}

//defer 用于释放锁，close链接，等
//执行时机： return 不是原子，step1，ret=返回值 step2: 真正的RET命令.
// defer在step1和step2之间执行

func main() {
	defer f(1)
	defer f(2)
	defer f(3) //先进后出
	f(0)

	fmt.Printf("%v\n", f1())
	fmt.Println()
	fmt.Printf("%v\n", f2())
	fmt.Println()
	fmt.Printf("%v\n", f4())
	fmt.Println()
	fmt.Printf("%v\n", f3())
	fmt.Println()

}
