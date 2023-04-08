/*package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int64
}

func incrementarEdad(p *Persona) {
	p.edad++
}

func buscarEdad(personas [5]Persona, edad int) *Persona {
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad {
			return &personas[i]
		}
	}
	return nil
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) *Persona {
	p := Persona{nombre, edad, ciudad, telefono}
	return &p
}

func actualizarTelefono(p *Persona, nuevoTelefono int64) {
	p.telefono = nuevoTelefono
}

func main() {
	// Crear dos Personas
	persona1 := Persona{"lucas", 31, "Comodoro Rivadavia", 87654533}
	persona2 := Persona{"misa", 3, "Caleta", 12234566}

	// Incrementar la edad de persona1
	incrementarEdad(&persona1)

	// Buscar una persona por edad
	personas := [5]Persona{persona1, persona2}
	edadBuscada := 30
	personaEncontrada := buscarEdad(personas, edadBuscada)
	if personaEncontrada != nil {
		fmt.Printf("Se encontro a la persona con edad %d: %v\n", edadBuscada, *personaEncontrada)
	} else {
		fmt.Printf("No se encontro a ninguna persona con edad %d\n", edadBuscada)
	}

	edadNoExistente := 40
	personaNoEncontrada := buscarEdad(personas, edadNoExistente)
	if personaNoEncontrada != nil {
		fmt.Printf("Se encontro a la persona con edad %d: %v\n", edadNoExistente, *personaNoEncontrada)
	} else {
		fmt.Printf("No se encontro a ninguna persona con edad %d\n", edadNoExistente)
	}

	// Crear una nueva persona y actualizar su teléfono
	personaNueva := crearPersona("paula", 28, "Buenos Aires", 12345678)
	fmt.Printf("Antes de actualizar el teléfono: %v\n", *personaNueva)
	actualizarTelefono(personaNueva, 21232131)
	fmt.Printf("Después de actualizar el teléfono: %v\n", *personaNueva)
}
*/