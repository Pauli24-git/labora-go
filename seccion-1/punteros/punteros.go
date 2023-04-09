package main

import (
	"fmt"
)

func main() {

	var a = 10
	var b = 20
	var ptrA = &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	a = b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	//a, b = b, a

	//fmt.Println(a)
	//fmt.Println(b)
}
