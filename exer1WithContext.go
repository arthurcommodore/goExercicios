package main

import (
	"context"
	"sync"
	"fmt"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

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
			fmt.Println("pong cancelado")
		case <-time.After(1 * time.Second):
			fmt.Println("pong")
	}
}

func main() {
	var wg sync.WaitGroup

	ctx,cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		<-time.After(1 * time.Second)
		cancel()
	}()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	wg.Wait()
	fmt.Println("Terminado workers") 
}
