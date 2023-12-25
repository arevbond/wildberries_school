/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации
*/

package main

import (
	"fmt"
	"strings"
)

// Пример кода:
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}
//

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println("Взялось 100 байт:", justString)

	runes := []rune(v)
	justString = string(runes[:100])
	fmt.Println("Взялось 100 символов:", justString)
}

func createHugeString(num int) string {
	var result strings.Builder
	for i := 0; i < num; i++ {
		fmt.Fprintf(&result, "界")
	}
	return result.String()
}

func main() {
	someFunc()
}
