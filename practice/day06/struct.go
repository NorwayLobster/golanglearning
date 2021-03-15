package main

import "fmt"

type myInt int     //自定义类型
type yourInt = int //类型别名
// byte, rune  也是类型别名  //编译时直接替换？？？

type Hello struct {
	a1 string
	a2 string
}
type Student struct { //自定义类型{//
	age    int
	name   string
	weight float32
	hobby  []string
	ptr    *Student
	Hello  //匿名成员
	// hello  Hello
	hi struct {
		h1 string
		h2 string
	}
}

func main() {
	var a1 myInt
	fmt.Printf("a1:%T\n", a1)
	var a2 yourInt
	fmt.Printf("a2:%T\n", a2)
	var a3 byte
	fmt.Printf("a3:%T\n", a3)
	var a4 rune
	a4 = '在' // wrong: "在"
	fmt.Printf("a4:%T\n", a4)
	var s1 Student
	fmt.Printf("s1:%T\n", s1)
	fmt.Printf("s1:%v\n", s1)
	s1.age = 10
	s1.hobby = []string{"篮球", "足球"}
	// s1.hello.a1 = "world"
	s1.a1 = "world"
	fmt.Printf("s1:%v\n", s1)

	var hi struct { ///匿名结构体
		h1 string
		h2 string
	}
	fmt.Printf("hi:%v\n", hi)

	ptr1 := &s1
	(*ptr1).age = 1
	fmt.Printf("s1:%v\n", s1)
	ptr1.age = 100 //语法糖
	fmt.Printf("s1:%v\n", s1)

	var s2 Student
	fmt.Printf("s2:%v\n", s2)
	ptr2 := &(s2.age) //成员变量地址
	*ptr2 = 100
	fmt.Printf("ptr2:%T\n", ptr2)
	fmt.Printf("s2:%v\n", s2)

	fmt.Println()
	//init: 值列表
	var hello1 = Hello{ //
		"world",
		"China",
	}
	fmt.Printf("hello1:%T\n", hello1)
	fmt.Printf("hello1:%v\n", hello1)
	var ptr_hello1 = &Hello{ //
		"world",
		"China",
	}
	fmt.Printf("ptr_hello1:%T\n", ptr_hello1)
	fmt.Printf("ptr_hello1:%v\n", ptr_hello1)
	//init: key-value
	var hello2 = Hello{ //
		a2: "world",
		a1: "China",
	}
	fmt.Printf("hello2:%T\n", hello2)
	fmt.Printf("hello2:%v\n", hello2)

	var ptr3 = new(Student)
	(*ptr3).name = "cheng"
	ptr3.name = "cheng"
	fmt.Printf("ptr3:%v\n", ptr3)

	// var s1_student = make(Student, 2, 10) //error
	// fmt.Println(s1_student)
}
