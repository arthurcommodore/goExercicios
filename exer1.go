package main

import (
	"sync"
	"fmt"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	ping()
	pong()
}

func ping() {
	fmt.Println("ping")
	time.Sleep(1 * time.Second)
}

func pong() {
	fmt.Println("pong")
	time.Sleep(1 * time.Second)
}


func main() {
	var wg sync.WaitGroup;

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("Todos os workers terminados...")
}
