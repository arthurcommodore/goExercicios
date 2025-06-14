package main

import (
	"fmt"
	"sync"
)

type ContadorSeguro struct {
	count int
	mut sync.Mutex
}

func (c *ContadorSeguro) Incrementar() {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.count++
}

func (c *ContadorSeguro) Valor() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.count
}


func main() {
	wg := sync.WaitGroup{}
	cts := ContadorSeguro{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cts.Incrementar()
		}()
	}

	wg.Wait()
	fmt.Println("Fim do programa, contador: ", cts.Valor())

}
