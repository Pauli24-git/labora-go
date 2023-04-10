package main

import "fmt"

func main() {

	peliculas := [10]string{"Shrek", "Matrix", "El Resplandor", "Harry Potter", "Cars", "Titanic", "Robocop", "Terminator", "Madagascar", "Parasite"}

	fmt.Println(peliculas)
	fmt.Println(peliculas[1])
	fmt.Println(len(peliculas))

	numeros := []int{1, 2, 3, 4, 5}
	longitud := len(numeros)
	fmt.Println(longitud)
}
