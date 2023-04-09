package main

import "fmt"

func main() {

	var paula = Persona{"Paula", 28, "Buenos Aires", 44441111}
	var lucas = Persona{Nombre: "Lucas", Edad: 28, Telefono: 19923451, Ciudad: "Comodoro Rivadavia"}

	fmt.Println(paula, lucas)
	cambiarCiudad(&lucas, "Lima")
	fmt.Println(lucas.Ciudad)
	cambiarCiudad(&paula, "Buenos Aires")
	fmt.Println(paula.Ciudad)
	fmt.Println(paula, lucas)

}

type Persona struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Telefono int
}

func cambiarCiudad(persona *Persona, ciudad string) {
	if persona.Ciudad != ciudad {
		persona.Ciudad = ciudad
	}
}
