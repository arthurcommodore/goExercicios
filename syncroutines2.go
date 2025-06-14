package main

import (
	"fmt"
	"sync"
)


func main() {

	var n int
	ch := make(chan int)

	fmt.Println("Digite n:")
	fmt.Scanln(&n)
	
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	
	for i := 0; i < n; i++ {
		fmt.Println(<-ch)
	}
}
