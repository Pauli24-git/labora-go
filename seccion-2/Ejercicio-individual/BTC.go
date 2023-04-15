package main

import (
	"fmt"
	"strings"
)

func main() {
	usuarios := [10]string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}

	monedasDisponibles := 50
	balanceUsuarios := make(map[string]int)

	monedasValores := make(map[string]int)
	monedasValores["a"] = 1
	monedasValores["e"] = 1
	monedasValores["i"] = 2
	monedasValores["o"] = 3
	monedasValores["u"] = 4
	monedasValores["A"] = 1
	monedasValores["E"] = 1
	monedasValores["I"] = 2
	monedasValores["O"] = 3
	monedasValores["U"] = 4

	balanceUsuarios, monedasDisponibles = ObtenerBalance(usuarios, monedasDisponibles, monedasValores)
	fmt.Println(balanceUsuarios)
	fmt.Println("Monedas restantes: ", monedasDisponibles)
}

func ObtenerBalance(usuarios [10]string, monedas int, valores map[string]int) (map[string]int, int) {
	balance := make(map[string]int)

	for _, u := range usuarios {
		balance[u], monedas = ValorizarVocales(u, monedas, valores)
	}

	return balance, monedas
}

func ValorizarVocales(nombre string, monedasRestantes int, valores map[string]int) (int, int) {
	var valorVocales int
	var monedasParaAgregar int
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		_, encontrado := valores[letra]

		if encontrado && valores[letra] <= monedasRestantes {

			monedasParaAgregar = valorVocales + valores[letra]

			if monedasParaAgregar >= 10 {
				monedasRestantes -= 10 - valorVocales
				valorVocales = 10
			} else {
				valorVocales += valores[letra]
				monedasRestantes -= valores[letra]
			}
		}
	}
	return valorVocales, monedasRestantes
}
