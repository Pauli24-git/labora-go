// Mostrar el nombre de los planetas del sistema solar, utilizando condiciones, structs y repetidores, mostrando su cantidad de lunas correctamente.
package main

import "fmt"

func main() {

	var planeta1 = planetas{"Mercurio", 0}
	var planeta2 = planetas{"Venus", 0}
	var planeta3 = planetas{"Tierra", 1}
	var planeta4 = planetas{"Marte", 2}
	var planeta5 = planetas{"Jupiter", 79}
	var planeta6 = planetas{"Saturno", 53}
	var planeta7 = planetas{"Urano", 27}

	sistemaSolar := [7]planetas{planeta1, planeta2, planeta3, planeta4, planeta5, planeta6, planeta7}

	for i := 0; i < 7; i++ {
		fmt.Printf("El planeta %s tiene %d lunas\n", sistemaSolar[i].Nombre, sistemaSolar[i].Lunas)
	}
}

type planetas struct {
	Nombre string
	Lunas  int
}
