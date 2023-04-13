package main

import (
	"fmt"
	"strings"
)

func main() {
	prueba := "Hola buenos dias. Buenas Mañanas. Buenas Noches"

	mapa := contadorPalabras(prueba)

	fmt.Println(mapa)
}

func contadorPalabras(frase string) map[string]int {
	v := strings.Fields(frase)

	m := make(map[string]int)
	for _, palabra := range v {
		_, encontrado := m[palabra]

		if encontrado {
			m[palabra] = m[palabra] + 1
		} else {
			m[palabra] = 1
		}
	}

	return m
}
