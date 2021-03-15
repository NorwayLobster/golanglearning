package main

import "fmt"

func main() {
	var age int = 20
	if age := 19; age >= 19 {
		fmt.Printf("age:%d\n", age)
	} else if age >= 10 {
		fmt.Printf("teenage, age:%d\n", age)
	} else {
		fmt.Printf("kid,age:%d\n", age)
	}

	fmt.Printf("age:%d\n", age)

}
