package main

import (
	"fmt"
)

func main() {

	var a = 10
	var b = 20

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a)
	fmt.Println(b)

}
