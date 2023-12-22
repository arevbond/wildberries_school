/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
)

func squares(arr []int) <-chan int {
	c := make(chan int)
	for _, n := range arr {
		go func(num int) {
			fmt.Println("Положил в канал:", num*num)
			c <- num * num
		}(n)
	}
	return c
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	c := squares(arr)
	result := 0
	for i := 0; i < len(arr); i++ {
		value := <-c
		result += value
	}
	fmt.Println(result)
}
