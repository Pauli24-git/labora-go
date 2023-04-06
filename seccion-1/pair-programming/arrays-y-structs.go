package main

import "fmt"

func main() {

	var paula = Persona{"Paula", 28, "Buenos Aires", 44441111}
	//var lucas = Persona{Nombre: "Lucas", Edad: 28, Telefono: 19923451, Ciudad: "Comodoro Rivadavia"}
	//var hugo = Persona{"Hugo", 63, "La Plata", 12345678}
	//var cristina = Persona{"Cristina", 60, "Bariloche", 87654321}
	//var misa = Persona{"Misa", 3, "El Condado", 11223344}

	incrementarEdad(&paula)
	buscarEdad()
}

type Persona struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Telefono int64
}

func incrementarEdad(persona *Persona) {
	persona.Edad++
	fmt.Println(persona.Edad)
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

/*

func actualizarTelefono(p *Persona, nuevoTelefono int64) {
    p.telefono = nuevoTelefono
}

func main() {
    // Crear dos Personas
    persona1 := Persona{"Juan", 25, "Madrid", 123456789}
    persona2 := Persona{"María", 30, "Barcelona", 987654321}

    // Incrementar la edad de persona1
    incrementarEdad(&persona1)

    // Buscar una persona por edad
    personas := [5]Persona{persona1, persona2}
    edadBuscada := 30
    personaEncontrada := buscarEdad(personas, edadBuscada)
    if personaEncontrada != nil {
        fmt.Printf("Se ha encontrado a la persona con edad %d: %v\n", edadBuscada, *personaEncontrada)
    } else {
        fmt.Printf("No se ha encontrado a ninguna persona con edad %d\n", edadBuscada)
    }

    edadNoExistente := 40
    personaNoEncontrada := buscarEdad(personas, edadNoExistente)
    if personaNoEncontrada != nil {
        fmt.Printf("Se ha encontrado a la persona con edad %d: %v\n", edadNoExistente, *personaNoEncontrada)
    } else {
        fmt.Printf("No se ha encontrado a ninguna persona con edad %d\n", edadNoExistente)
    }

    // Crear una nueva persona y actualizar su teléfono
    personaNueva := crearPersona("Pedro", 20, "Sevilla", 111111111)
    fmt.Printf("Antes de actualizar el teléfono: %v\n", *personaNueva)
    actualizarTelefono(personaNueva, 222222222)
    fmt.Printf("Después de actualizar el teléfono: %v\n", *personaNueva)
}
*/
