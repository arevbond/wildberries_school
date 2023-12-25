/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func New(x, y int) Point {
	return Point{x: x, y: y}
}

func findDistance(a, b Point) float64 {
	result := math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	a := New(10, 12)
	b := New(0, 0)
	dist := findDistance(a, b)
	fmt.Printf("Расстрояние между {%d %d} и {%d %d} = %.2f\n",
		a.x, a.y, b.x, b.y, dist)
}
