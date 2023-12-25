/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»
*/

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Fields(s)
	l, r := 0, len(strs)-1
	for l < r {
		strs[l], strs[r] = strs[r], strs[l]
		l++
		r--
	}
	return strings.Join(strs, " ")
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
	fmt.Println(reverseWords("третье второе Первое"))
	fmt.Println(reverseWords("ВВ бб   АА"))
}
