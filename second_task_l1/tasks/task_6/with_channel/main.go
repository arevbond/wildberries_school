/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Способ 1. С помощью канала

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(quit <-chan bool) {
	for {
		select {
		default:
			fmt.Println("worker do something...")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		case <-quit:
			fmt.Println("worker stopped")
			return
		}
	}
}

func main() {
	quit := make(chan bool)

	go worker(quit)

	time.Sleep(5 * time.Second)
	quit <- true
}
