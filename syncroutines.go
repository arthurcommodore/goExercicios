package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	var n int

	fmt.Print("Digite n: ")
	fmt.Scanln(&n)

	// Criamos uma slice para armazenar os resultados
	resultados := make([]int, n)

	// Cada goroutine calcula e envia o valor junto com o índice
	for i := 0; i < n; i++ {
		go func(ind int) {
			ch <- ind // manda o índice (ou poderia ser resultado de alguma conta)
		}(i)
	}

	// Lemos n valores e armazenamos no índice correto
	for i := 0; i < n; i++ {
		valor := <-ch
		resultados[valor] = valor
	}

	// Agora imprimimos em ordem
	fmt.Println("Saída ordenada:")
	for _, v := range resultados {
		fmt.Println(v)
	}
}

