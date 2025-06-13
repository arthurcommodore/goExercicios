package main

import (
	"context"
	"sync"
	"fmt"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	select {
		case <-ctx.Done():
			fmt.Println("Terminado o tempo")
		default: 
	}

	ping(ctx)
	pong(ctx)
}


func ping(ctx context.Context) {
	select {
		case <-ctx.Done():
			fmt.Println("ping cancelado")
		case <-time.After(1 * time.Second):
			fmt.Println("ping")
	}
}

func pong(ctx context.Context) {
	select {
		case <-ctx.Done():
			fmt.Println("ping cancelado")
		case <-time.After(1 * time.Second):
			fmt.Println("ping")
	}
}

func main() {
	var wg sync.WaitGroup

	ctx,cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	wg.Wait()
	fmt.Println("Terminado workers") 
}
