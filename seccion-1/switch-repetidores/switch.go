package main

import "fmt"

func main() {
	var numeroIngresado int
	fmt.Println("Ingrese el numero del dia de la semana: ")
	fmt.Scan(&numeroIngresado)

	switch {
	case numeroIngresado == 1:
		fmt.Println("LUNES")
	case numeroIngresado == 2:
		fmt.Println("MARTES")
	case numeroIngresado == 3:
		fmt.Println("MIERCOLES")
	case numeroIngresado == 4:
		fmt.Println("JUEVES")
	case numeroIngresado == 5:
		fmt.Println("VIERNES")
	case numeroIngresado == 6:
		fmt.Println("SABADO")
	case numeroIngresado == 7:
		fmt.Println("DOMINGO")
	default:
		fmt.Println("El numero que ingresó no corresponde a ningun dia de la semana")
	}
}
