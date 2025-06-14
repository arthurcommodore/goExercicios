package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	fmt.Println("Digite n")
	var n int
	fmt.Scanln(&n)
	
	for i := 0; i < n; i++ {
		go func(v int) {
			ch <- v
		}(i)
	}
	
	for i := 0; i < n; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}
