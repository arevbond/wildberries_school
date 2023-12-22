/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("worker 1: положил %d в канал\n", n)
			out <- n
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(out)
	}()
	return out
}

func square(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range nums {
			fmt.Printf("worker 2: положил %d в канал\n", n*n)
			out <- n * n
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(out)
	}()
	return out
}

func main() {
	c := generate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	out := square(c)
	for n := range out {
		fmt.Println("Результат на выходе:", n)
	}
}
