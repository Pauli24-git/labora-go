package main

import (
	"fmt"
	"sort"
	"strconv"
)

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
			ordenarEstudiantes(Estudiantes)
		case numeroIngresado == 3:
			buscarEstudiante(Estudiantes)
		}
		fmt.Println("Ingrese la accion que desea realizar:\n 1) Crear estudiante\n 2) Ordenar la lista\n 3) Buscar estudiante\n 4) Salir del programa\n")
		fmt.Scan(&numeroIngresado)
	}
	defer func(estudiantes []Estudiante) {
		var (
			suma     float64
			promedio float64
		)
		//		fmt.Println(len(estudiantes))
		for i := 0; i < len(estudiantes); i++ {
			suma = estudiantes[i].Nota + suma
		}
		//		fmt.Println(suma)
		promedio = suma / float64(len(estudiantes))
		fmt.Printf("El promedio de las notas es %.2f ", promedio)
	}(Estudiantes)

	defer func(estudiantes []Estudiante) {
		var (
			maxnota   float64
			minnota   float64
			alumnomax Estudiante
			alumnomin Estudiante
		)
		minnota = estudiantes[0].Nota
		alumnomin = estudiantes[0]
		for i := 0; i < len(estudiantes); i++ {
			if estudiantes[i].Nota > maxnota {
				maxnota = estudiantes[i].Nota
				alumnomax = estudiantes[i]
			}
			if estudiantes[i].Nota < minnota {
				minnota = estudiantes[i].Nota
				alumnomin = estudiantes[i]
			}

		}
		fmt.Printf("El estudiante con mayor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n", alumnomax.Nombre, alumnomax.Nota, alumnomax.Codigo)
		fmt.Printf("El estudiante con menor nota es:\n\t Nombre: %s, Nota: %.2f, Codigo: %s\n\n", alumnomin.Nombre, alumnomin.Nota, alumnomin.Codigo)

	}(Estudiantes)

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

func ordenarEstudiantes(listaOriginal []Estudiante) {
	var opcionEstudiante int
	var opcionOrden int

	fmt.Println("Indique preferencia al ordenar:\n 1) Segun Nombre \n2)Segun Nota \n 3)Segun Codigo")
	fmt.Scan(&opcionEstudiante)
	fmt.Println("Indique orden 1) Ascendente \n 2) Descendente")
	fmt.Scan(&opcionOrden)

	OrdenarSegunCriterio(listaOriginal, opcionEstudiante, opcionOrden)

}

func OrdenarSegunCriterio(estudiantes []Estudiante, ordenamiento int, tipoOrdenamiento int) {

	fmt.Println(estudiantes)

	//Aca probe con sort.Slice para recorrer y ordenar segun criterio. Esta manera la encontre despues de subir el ejercicio, pero no entendi muy bien como funciona el sort...

	if tipoOrdenamiento == 1 { //Por nombre
		if ordenamiento == 1 {
			sort.Slice(estudiantes, func(i, j int) bool {
				return estudiantes[i].Nombre < estudiantes[j].Nombre
			})

		} else {
			sort.Slice(estudiantes, func(i, j int) bool {
				return estudiantes[i].Nombre > estudiantes[j].Nombre
			})
		}
	}

	fmt.Println(estudiantes)

	if tipoOrdenamiento == 2 { //Por nota
		for i := 0; i < len(estudiantes)-1; i++ {

			for j := 0; j < len(estudiantes)-1-i; j++ {
				if ordenamiento == 1 {
					if estudiantes[j].Nota > estudiantes[j+1].Nota {
						estudiantes[j], estudiantes[j+1] = estudiantes[j+1], estudiantes[j]
					}
				} else {
					if estudiantes[j].Nota < estudiantes[j+1].Nota {
						estudiantes[j], estudiantes[j+1] = estudiantes[j+1], estudiantes[j]
					}
				}
			}
		}
	}

	if tipoOrdenamiento == 3 { //Por codigo
		for i := 0; i < len(estudiantes)-1; i++ {

			for j := 0; j < len(estudiantes)-1-i; j++ {
				if ordenamiento == 1 {
					if estudiantes[j].Codigo > estudiantes[j+1].Codigo {
						estudiantes[j], estudiantes[j+1] = estudiantes[j+1], estudiantes[j]
					}
				} else {
					if estudiantes[j].Codigo < estudiantes[j+1].Codigo {
						estudiantes[j], estudiantes[j+1] = estudiantes[j+1], estudiantes[j]
					}
				}
			}
		}
	}
}

func buscarEstudiante(listaOriginal []Estudiante) {
	var opcionAtributo string
	var opcionValor string

	fmt.Println("Indique cual atributo desea buscar:\n 1) Segun Nombre \n2)Segun Nota \n 3)Segun Codigo")
	fmt.Scan(&opcionAtributo)
	fmt.Println("Ingrese el valor:\n ")
	fmt.Scan(&opcionValor)

	var opcionValorConvertida int64

	opcionValorConvertida, _ = strconv.ParseInt(opcionAtributo, 10, 64)

	if opcionValorConvertida == 1 {
		for i := 0; i < len(listaOriginal); i++ {
			if opcionValor == listaOriginal[i].Nombre {
				fmt.Println("Usuario encontrado\n")
				fmt.Println(listaOriginal[i])
			}
		}
	}

	if opcionValorConvertida == 2 {
		for i := 0; i < len(listaOriginal); i++ {
			var opcionConvertida float64

			opcionConvertida, _ = strconv.ParseFloat(opcionValor, 64)

			if opcionConvertida == listaOriginal[i].Nota {
				fmt.Println("Usuario encontrado\n")
				fmt.Println(listaOriginal[i])
			}
		}
	}

	if opcionValorConvertida == 3 {
		for i := 0; i < len(listaOriginal); i++ {
			if opcionValor == listaOriginal[i].Codigo {
				fmt.Println("Usuario encontrado\n")
				fmt.Println(listaOriginal[i])
			}
		}
	}

}
