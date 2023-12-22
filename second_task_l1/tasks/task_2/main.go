/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func printSquares(num int) {
	defer wg.Done()
	fmt.Println(num * num)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	for _, n := range arr {
		wg.Add(1)
		go printSquares(n)
	}
	wg.Wait()
}
