package main

import (
	"fmt"
	"sync"
)

var contador int
var wg sync.WaitGroup

func incrementar() {
	for i := 0; i < 10000; i++ {
		contador++
	}
	wg.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementar()
	}

	wg.Wait()
	fmt.Println("Contador final:", contador)
}

