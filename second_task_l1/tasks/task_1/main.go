/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

// https://go.dev/doc/effective_go#embedding

package main

import "fmt"

type Some interface {
}

// first example
type Human struct {
	name string
	age  int
}

func (h *Human) Age() int {
	return h.age
}

type Action struct {
	Human
}

func (a Action) GetAge() int {
	return a.Age()
}

// end first example

// second example
type Human2 struct {
	name string
	age  int
}

func (h *Human2) Age() int {
	return h.age
}

type Action2 struct {
	h Human
}

func (a Action2) GetAge() int {
	return a.h.Age()
}

// end second example
func main() {
	a := Action{Human{age: 21, name: "nikita"}}
	fmt.Println(a.name)
	fmt.Println(a.Age())
}
