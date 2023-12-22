/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/

package main

import "fmt"

// changeBit устанавливает i-ый бит числа num в противположное значение.
func changeBit(num int64, i int) int64 {
	fmt.Printf("Пришло:\n%064b\nНеобходимо переставить %d бит\n",
		num, i)
	num ^= 1 << i
	fmt.Printf("Результат:\n%064b\n", num)
	return num
}

func main() {
	var num int64 = 8
	fmt.Printf("Начальное число: %d\n", num)
	result := changeBit(num, 3)
	fmt.Println("Число после изменения бита:", result)
}
