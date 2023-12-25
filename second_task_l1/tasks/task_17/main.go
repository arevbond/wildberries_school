/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import "fmt"

func binSearch(arr []int, target int) (int, bool) {
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) / 2
		if arr[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if arr[l] == target {
		return l, true
	}
	return -1, false
}

func main() {
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println(binSearch([]int{1, 2, 3, 4, 6, 7}, 5))
}
