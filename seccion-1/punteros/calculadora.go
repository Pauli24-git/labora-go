package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	fmt.Scan(&num1, &num2)
	fmt.Println(num1, num2)

	fmt.Println("Ingrese el valor numero 1: ", num1)
	fmt.Println("Ingrese el valor numero 2: ", num2)

	calcular(&num1, &num2)

}

func calcular(num1, num2 *int) {
	var suma = *num1 + *num2
	fmt.Println("Suma: ", suma)
	var resta = *num1 - *num2
	fmt.Println("Resta: ", resta)
	var multiplicacion = *num1 * *num2
	fmt.Println("Multi: ", multiplicacion)
	var division = *num1 / *num2
	fmt.Println("Division:", division)
}
