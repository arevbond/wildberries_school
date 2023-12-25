/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func bigIntSum(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	return c.Add(a, b)
}

func bigIntSub(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	return c.Sub(a, b)
}

func bigIntMul(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	return c.Mul(a, b)
}

func bigIntDiv(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	return c.Div(a, b)
}

func main() {
	a, b := big.NewInt(1e13), big.NewInt(2e11)
	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Println("Сумма a+b:", bigIntSum(a, b))
	fmt.Println("Разность a-b:", bigIntSub(a, b))
	fmt.Println("Произведение a*b:", bigIntMul(a, b))
	fmt.Println("Частное a/b:", bigIntDiv(a, b))
}
