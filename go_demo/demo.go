package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", len(os.Args))
		fmt.Println("hello world", os.Args[0])
		fmt.Println("hello world", os.Args[1])
		fmt.Println("hello world", os.Args[2])
	}
	os.Exit(-1)
}
