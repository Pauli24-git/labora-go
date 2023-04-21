package main

import (
	"fmt"
)

func multiplicarFilas(filaA []int, matrizB [][]int, result chan []int) {
	n := len(matrizB[0])
	m := len(filaA)
	resultadoFila := make([]int, n)
	for j := 0; j < n; j++ {
		sum := 0
		for k := 0; k < m; k++ {
			sum += filaA[k] * matrizB[k][j]
		}
		resultadoFila[j] = sum
	}
	result <- resultadoFila
}

func multiplicarMatrices(matrizA [][]int, matrizB [][]int) [][]int {
	m := len(matrizA)
	n := len(matrizB[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	resultChan := make(chan []int, m)
	for i := 0; i < m; i++ {
		go multiplicarFilas(matrizA[i], matrizB, resultChan)
	}

	for i := 0; i < m; i++ {
		resultadoFila := <-resultChan
		result[i] = resultadoFila
	}

	close(resultChan)
	return result
}

func main() {
	matrizA := [][]int{
		{3, 2, 1},
		{4, 5, 6},
	}

	matrizB := [][]int{
		{12, 11},
		{10, 9},
		{8, 7},
	}

	result := multiplicarMatrices(matrizA, matrizB)
	fmt.Println(result)
}
