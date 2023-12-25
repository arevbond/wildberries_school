/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse("главрыба"))
	fmt.Println(reverse("abcd"))
	fmt.Println(reverse(""))
	fmt.Println(reverse("东丝丟丟丞"))
}
