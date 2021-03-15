package main

import "fmt"

func main() {
	fmt.Println("人生苦短，let's Go!")
	fmt.Print("hello world!\t")
	fmt.Println("hello world!\r")
	var str string = "Hello Go"
	fmt.Printf("%s!\n", str)
	fmt.Printf("%#v!\n", str)

	var ch1 byte = 'h'
	ch2 := 'h'
	ch3 := '中'
	fmt.Printf("ch1:%c\n", ch1)
	fmt.Printf("ch1:%T\n", ch1)
	fmt.Printf("ch2:%c\n", ch2)
	fmt.Printf("ch2:%T\n", ch2)
	fmt.Printf("ch3:%c\n", ch3)
	fmt.Printf("ch3:%T\n", ch3)

	var i1 = 101
	var i2 = 0x1234
	var i3 = 017
	var i4 = int8(17) //明确指定类型，否则默认为int
	fmt.Printf("%d\n", i1)
	fmt.Printf("%x\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%b\n", i1)

	fmt.Println()

	fmt.Printf("%d\n", i2)
	fmt.Printf("%x\n", i2)
	fmt.Printf("%o\n", i2)
	fmt.Printf("%b\n", i2)
	fmt.Println()

	fmt.Printf("%d\n", i3)
	fmt.Printf("%x\n", i3)
	fmt.Printf("%o\n", i3)
	fmt.Printf("%b\n", i3)

	fmt.Println()
	fmt.Printf("%T\n", i4)
	fmt.Printf("i4:%v\n", i4)  //%v 万能
	fmt.Printf("i4:%#v\n", i4) //%v 万能
}
