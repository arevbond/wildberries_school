/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"unicode"
)

func checkUnique(s string) bool {
	set := make(map[rune]struct{})
	runes := []rune(s)
	for _, r := range runes {
		r = unicode.ToLower(r)
		if _, ok := set[r]; !ok {
			set[r] = struct{}{}
		}
	}
	return len(set) == len(runes)
}

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))

	fmt.Println(checkUnique("АаББ"))
	fmt.Println(checkUnique("АБвгЖ"))
	fmt.Println(checkUnique(""))
}
