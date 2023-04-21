package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	go func() {
		c <- sumaNumeros()
	}()

	x := <-c

	fmt.Println(x)
}

func sumaNumeros() int {
	s := 10
	for i := 0; i < 100; i++ {
		s += i
	}
	return s
}

// Ejercicio 1: Suma de números en un rango utilizando gorutinas y canales
// Escribe un programa en Go que sume todos los números en un rango dado (por ejemplo, de 1 a 100)
// utilizando gorutinas y canales para dividir el trabajo entre varias gorutinas.
