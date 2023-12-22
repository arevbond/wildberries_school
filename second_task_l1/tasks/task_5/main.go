/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channel(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%d) %s", i, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	timeout := time.After(5 * time.Second)
	c1 := channel("some data from channel 1")
	c2 := channel("some data from channel 2")
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("Время программы вышло")
			return
		}
	}
}
