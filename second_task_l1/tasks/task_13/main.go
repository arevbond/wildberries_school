/*
Поменять местами два числа без создания временной переменной.
*/

package main

import "fmt"

func main() {
	number1, number2 := 10, 20
	fmt.Printf("Number1: %d Number2: %d\n", number1, number2)

	number1, number2 = number2, number1
	fmt.Printf("Number1: %d Number2: %d\n", number1, number2)
}
