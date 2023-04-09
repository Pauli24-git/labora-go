package main

import (
	"fmt"
)

func main() {

	var a = 10
	var b = 20

	var ptrA *int

	ptrA = &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ptrA)
	fmt.Println(*ptrA)
	fmt.Println(&a)

	a = b
	b = *ptrA

	fmt.Println(a)
	fmt.Println(b)

}
