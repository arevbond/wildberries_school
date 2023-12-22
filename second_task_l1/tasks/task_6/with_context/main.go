/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Способ 2. С помощью context

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done work")
			return
		default:
			fmt.Println("worker do some work...")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	wg.Wait()
}
