package main

import "fmt"

func main() {
	restaurante := "McDonalds"
	Stars := 2.3
	fmt.Printf("%s tiene una calificacion de %.1f estrellas\n", restaurante, Stars)

	restaurante = "Dominos"
	Stars = 4.5
	fmt.Printf("%s tiene una calificacion de %.1f estrellas\n", restaurante, Stars)

	restaurante = "Uggis"
	Stars = 1
	fmt.Printf("%s tiene una calificacion de %.1f estrellas\n", restaurante, Stars)

}
