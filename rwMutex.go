package main

import (
	"fmt"
	"sync"
	"time"
)

type ContadorSeguro struct {
	count int
	mu    sync.RWMutex
}

func (c *ContadorSeguro) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *ContadorSeguro) LerValor() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	contador := &ContadorSeguro{}

	// Escritores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				contador.Incrementar()
			}
		}()
	}

	// Leitores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				_ = contador.LerValor() // só lendo
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Valor final:", contador.LerValor())
}

