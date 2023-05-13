package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	var wg sync.WaitGroup
	//var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; counter < 200; j++ {
				//		mutex.Lock()
				counter++
				//		mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter value:", counter)
}
