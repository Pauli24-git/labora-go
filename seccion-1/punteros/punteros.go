package main

import (
	"fmt"
)

func main() {

	var a = 10
	var b = 20

	fmt.Println(a)
	fmt.Println(b)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
