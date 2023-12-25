/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"go/types"
	"reflect"
)

func checkType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("Переменная %d имеет типо int\n", value)
	case string:
		fmt.Printf("Переменная %s имеет типо string\n", value)
	case bool:
		fmt.Printf("Переменная %t имеет типо bool\n", value)
	case types.Chan:
		fmt.Println("Переменная value имеет типо chan")
	default:
		fmt.Println("Переменная value не является типом: " +
			"int, string, bool, chan")
	}
}

func checkTypeWithReflect(value interface{}) {
	t := reflect.TypeOf(value)
	fmt.Println("Переменная value имеет тип:", t.String())
}

func main() {
	//checkType(1)
	//checkType("1")
	//checkType(true)
	//checkType(false)
	//c := make(chan int)
	//checkType(c)
	//fmt.Println()

	checkTypeWithReflect(1)
	checkTypeWithReflect("1")
	checkTypeWithReflect(true)
	checkTypeWithReflect(false)
	checkTypeWithReflect(make(chan int))
}
