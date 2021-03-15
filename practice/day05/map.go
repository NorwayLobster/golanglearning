package main

import (
	"fmt"
	"sort"
)

func main() {
	//0.   申请空的hash并把其绑定到引用m0上
	m0 := map[string]int{}
	fmt.Println(m0)
	//1.
	var m2 map[string]int = map[string]int{
		"cheng":    30,
		"yangyang": 28,
	}
	fmt.Println(m2)

	//2.
	m3 := map[string]int{
		"cheng":    30,
		"yangyang": 28,
	}
	fmt.Println(m3)

	//3.
	var m1 map[string]int
	fmt.Println(m1 == nil)
	//必须init，即申请内存
	m1 = make(map[string]int, 10) //区别？
	// m1 = make(map[string]int) //会有扩容和rehash开销???
	fmt.Println(m1, len(m1)) //, cap(m1))

	m1["cheng"]++
	fmt.Println(m1, len(m1))
	fmt.Println(m1["cheng"])

	//4.
	m1["wang"] = 19
	m1["yangb"] = 18
	m1["yanga"] = 17
	name, ok := m1["cheng"]
	if !ok {
		fmt.Println("无此key", name)
	}

	for name, age := range m1 {
		fmt.Println(name, age)
	}

	//5.
	//get names
	var names []string
	for name := range m1 {
		names = append(names, name)
	}
	fmt.Println(names)
	//sort names
	sort.Strings(names)
	fmt.Println(names)
	//print age sortedly by name
	for i, name := range names {
		fmt.Println(i, name, m1[name])
	}

	//6.
	delete(m1, "wang")
	fmt.Println(m1, len(m1)) //, cap(m1))
	delete(m1, "hu")

	//7.map的数组的切片
	var s1 = make([]map[string]int, 2, 10) //
	fmt.Println(s1, len(s1))               //, cap(m1))
	s1[0] = make(map[string]int)           //绑定
	s1[1] = make(map[string]int, 10)       //绑定
	s1[0]["cheng"] = 30
	s1[0]["wang"] = 31
	s1[1]["hu"] = 19
	s1[1]["yang"] = 18
	fmt.Println(s1, len(s1))
	fmt.Println(s1[0], len(s1[0]))
	fmt.Println(s1[1], len(s1[1]))
	fmt.Printf("%T\n", s1)

}
