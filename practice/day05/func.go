package main

import "fmt"

//0.
func sum0(a int, b int) (ret int) { //多个返回值
	ret = a + b //命名返回值，可以直接使用
	return
}

//1.
func sum(a int, b int) (ret int, ret1 int) { //多个返回值
	ret1 = a + b //命名返回值，可以直接使用
	return a + b, b
}

//2.
func sum1(a, b, c int) (ret, ret1 int) { //紧凑写法
	return a + b, c
}

//3.
func sum2(a, b, c int, d ...int) (ret, ret1 int) { //变参数,0～n个皆可，
	fmt.Printf("%T\n", d) ///d为其底层数据类型的切片
	return a + b, c
}

//4.
func f1(a int) int { //无返回值
	fmt.Println(a)
	return a
}

//5.
func f2(a int) int { // 匿名返回值
	fmt.Println(a)
	return 3
}

//6. 参数是函数对象，返回值也是函数对象
func f3(a func(int) int) (b func(int) int) {
	a(1)
	fmt.Printf("a type:%T\n", a)
	return a
}
func main() {
	ret, ret1 := sum(1, 2)
	fmt.Printf("ret:%v\n", ret)
	fmt.Println(ret1)
	sum2(1, 2, 3)
	sum0(1, 2)
	sum(1, 2)
	sum1(1, 2, 3)
	f1(1)
	a1 := f1
	a1(1)
	fmt.Printf("a1 type:%T\n", a1)

	a2 := f3(a1)
	fmt.Printf("a2 type:%T\n", a2)

}
