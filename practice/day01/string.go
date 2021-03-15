package main

import (
	"fmt"
	"strings"
)

func main() {
	c1 := 'h'
	c2 := '中'
	fmt.Printf("%T\n", c1)
	fmt.Printf("c1:%c\n", c1)

	fmt.Println()

	fmt.Printf("%T\n", c2)
	fmt.Printf("c2:%v\n", c2)
	fmt.Printf("c2:%c\n", c2)
	fmt.Println()
	// "字符"
	var str1 string
	var str2 string = "hello world"
	fmt.Printf("%T\n", str1)
	fmt.Printf("str1:%v\n", str1)
	fmt.Println()
	fmt.Printf("%T\n", str2)
	fmt.Printf("str2:%v\n", str2)

	fmt.Println("\\ ' \"")

	fmt.Println()

	fmt.Println("I'm Okay")
	var str3 = `您好，
		你好
		hello
		`
	fmt.Print(str3)
	fmt.Printf("str3:%s\n", str3)
	fmt.Println()
	var str4 = str3 + str1
	fmt.Printf("str4:%s\n", str4)
	fmt.Println()
	var str5 = fmt.Sprintf("%s%s", str3, str4)
	fmt.Printf("str5:%s\n", str5)
	fmt.Println()
	var str6 string = "hello \\ world "
	ret := strings.Split(str6, "\\")
	fmt.Print(ret)
	fmt.Println()
	fmt.Println(strings.Contains(str2, "hello"))
	fmt.Println(strings.HasPrefix(str2, "hello"))
	fmt.Println(strings.HasSuffix(str2, "world"))
	fmt.Println(strings.Index(str2, "hello"))
	fmt.Println(strings.LastIndex(str2, "hello"))
	fmt.Println(strings.Join(ret, "+"))

}
