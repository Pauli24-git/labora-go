package main

import "fmt"

func main() {

	var Estudiantes []Estudiante

	var numeroIngresado int
	fmt.Println("Ingrese la accion que desea realizar:\n 1) Crear estudiante\n 2) Ordenar la lista\n 3) Buscar estudiante\n 4) Salir del programa\n")
	fmt.Scan(&numeroIngresado)

	for numeroIngresado != 4 {
		switch {
		case numeroIngresado == 1:
			Estudiantes = append(Estudiantes, crearNuevoEstudiante())
		case numeroIngresado == 2:
			fmt.Println("")
		case numeroIngresado == 3:
			fmt.Println("")
		}
		fmt.Println(Estudiantes)
		fmt.Println("Ingrese la accion que desea realizar:\n 1) Crear estudiante\n 2) Ordenar la lista\n 3) Buscar estudiante\n 4) Salir del programa\n")
		fmt.Scan(&numeroIngresado)
	}
}

type Estudiante struct {
	Nombre string
	Nota   float64
	Codigo string
}

func pedirDatos(nombre *string, nota *float64, codigo *string) {
	fmt.Printf("Ingrese su nombre: ")
	fmt.Scan(nombre)

	fmt.Printf("Ingrese su nota: ")
	fmt.Scan(nota)

	fmt.Printf("Ingrese su codigo: ")
	fmt.Scan(codigo)
}

func NewEstudiante(nombre string, nota float64, codigo string) Estudiante {
	return Estudiante{Nombre: nombre, Nota: nota, Codigo: codigo}
}

func crearNuevoEstudiante() Estudiante {
	var estudianteIngresado Estudiante

	var nombre string
	var nota float64
	var codigo string

	pedirDatos(&nombre, &nota, &codigo)

	estudianteIngresado = NewEstudiante(nombre, nota, codigo)
	return estudianteIngresado
}

func ordenarEstudiantes(listaOriginal []Estudiante) []Estudiante {
	var opcionEstudiante int
	var opcionOrden int

	fmt.Println("Indique preferencia al ordenar:\n 1) Segun Nombre \n2)Segun Nota \n 3)Segun Codigo")
	fmt.Scan(&opcionEstudiante)
	fmt.Println("Indique orden 1) Ascendente \n 2) Descendente")
	fmt.Scan(&opcionOrden)

	switch {
	case opcionEstudiante == 1:
		OrdenarPorNombre(&listaOriginal, opcionOrden)
	}

}

func OrdenarPorNombre(lista *[]Estudiante, ordenamiento int) {

}
