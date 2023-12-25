/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/

package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)

		arr = quickSort(arr, low, p-1)

		arr = quickSort(arr, p+1, high)

	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{5, 4, 3, 1, 0, -5, 12, 6, 8, 3, 5, 6}
	res := quickSortStart(arr)
	fmt.Printf("Начальный массив: %v\nМассив после сортировки: %v\n", arr, res)
}
