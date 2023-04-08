package main

import "fmt"

func main() {

	var persona1 = Persona{"Paula", 23, "Buenos Aires", 232555323}
	var persona2 = Persona{Nombre: "Lucas", Edad: 28, Telefono: 19923451, Ciudad: "Comodoro Rivadavia"}
	//var hugo = Persona{"Hugo", 63, "La Plata", 12345678}
	//var cristina = Persona{"Cristina", 60, "Bariloche", 87654321}
	//var misa = Persona{"Misa", 3, "El Condado", 11223344}

	var people = [5]Persona{persona1, persona2}

	incrementarEdad(&persona1)
	fmt.Println(persona1.Nombre, persona1.Ciudad)
	fmt.Println(buscarEdad(people, 18))

	fmt.Println(persona1.personaNueva("Cristina", 60, "Bariloche", 87654321))

	actualizarTelefono(&persona1, 12121212)

}

type Persona struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Telefono int64
}

func (p *Persona) personaNueva(nombre string, edad int, ciudad string, telefono int64) Persona {
	var personaNueva = Persona{nombre, edad, ciudad, telefono}
	return personaNueva
}

func incrementarEdad(persona *Persona) {
	persona.Edad++
	fmt.Println(persona.Edad)
}

func buscarEdad(personas [5]Persona, edad int) Persona {
	var personaAuxiliar = Persona{}

	for i := 0; i < len(personas); i++ {
		if personas[i].Edad == edad {
			personaAuxiliar = personas[i]
			break
		}
	}
	return personaAuxiliar
}

func actualizarTelefono(persona *Persona, telefono int64) {
	persona.Telefono = telefono
	fmt.Println(persona.Telefono)
}
