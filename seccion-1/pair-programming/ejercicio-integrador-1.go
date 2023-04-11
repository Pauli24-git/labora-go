package main

import "fmt"

func main() {
	//var people = [5]Persona{
	var nombre string
	var edad uint
	var altura uint
	var peso uint

	//fmt.Scan(&nombre, &edad, &altura, &peso)
	//fmt.Println(nombre, edad, altura, peso)

}

type Personass struct {
	Nombre string
	Edad   int
	Altura int
	Peso   int
}

func ingresarDatosPersona(nombre string, edad *uint, altura *uint, peso *uint) {
	var ingresoCorrecto bool
	fmt.Println("Ingrese su nombre, edad, altura y peso: ")
	for !ingresoCorrecto {

		fmt.Scanln(&nombre, &edad, &altura, &peso)

		if (nombre) != "" {
			ingresoCorrecto = true
		} else {
			fmt.Println("Nombre Invalido. Ingrese nuevamente: ")
		}

	}
}
