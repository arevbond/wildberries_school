/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

package main

import (
	"fmt"
	"sort"
)

func getIntervals(temps []float64) [][]float64 {
	if len(temps) == 0 {
		return nil
	}

	sort.Float64s(temps)

	cur := make([]float64, 0)
	result := make([][]float64, 0)

	start := temps[0]
	delta := 10.0
	for _, temp := range temps {
		if start+delta >= temp {
			cur = append(cur, temp)
		} else {
			result = append(result, cur)
			cur = []float64{temp}
			start = temp
		}
	}
	if len(cur) > 0 {
		result = append(result, cur)
	}
	return result
}

func main() {
	intervals := getIntervals([]float64{-35.4, -27.0, 13.0, 19.0, 15.5, 24.5,
		-21.0, 32.5})
	for _, interval := range intervals {
		fmt.Println(interval)
	}
}
