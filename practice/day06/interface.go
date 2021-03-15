package main

import "fmt"

type dog struct{}
type cat struct{}
type person struct{}

type shout_interface interface { //实现该接口中所有方法的类型都是该interface的对象
	shout() (int, int) //只要定义了shout()和move()方法的类型都是shout_interface类型的对象
	move()
}

func (d dog) shout() (a, b int) {
	fmt.Println("wang wang wang")
	return 1, 2
}
func (d dog) move() { fmt.Println("dog move") }
func (c cat) shout() (a, b int) {
	fmt.Println("miao miao miao")
	return 1, 2
}
func (c cat) move() { fmt.Println("cat move") }
func (p person) shout() (a, b int) {
	fmt.Println("a a a")
	return 1, 2
}
func (p person) move() { fmt.Println("person move") }

func f(x shout_interface) {
	x.shout()
	x.move()
}

func main() {

	var d dog
	var c cat
	var p person
	f(d)
	f(c)
	f(p)
}
