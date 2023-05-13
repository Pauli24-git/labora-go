package main

import (
	"fmt"
	"sync"
)

var balance = 1000 // saldo de la cuenta bancaria

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	fmt.Println(balance)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go deposito(&wg, &m)
		wg.Add(1)
		go retiro(&wg, &m)
	}
	wg.Wait()
	fmt.Println(balance)
}

func deposito(wg *sync.WaitGroup, m *sync.Mutex) {
	var x = 100
	m.Lock()
	balance = balance + x
	fmt.Printf("despues del deposito: %v\n", balance)
	m.Unlock()
	wg.Done()
}

func retiro(wg *sync.WaitGroup, m *sync.Mutex) {
	var y = 300
	m.Lock()
	defer m.Unlock()
	if balance > y {
		balance = balance - y
		fmt.Printf("despues del retiro: %v\n", balance)
	}
	wg.Done()
}
