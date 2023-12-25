/*
Удалить i-ый элемент из слайса.
*/

package main

import "fmt"

func removeElement(arr []int, indx int) []int {
	result := []int{}
	result = append(result, arr[:indx]...)
	result = append(result, arr[indx+1:]...)
	return result
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	newArr := removeElement(arr, 0)
	fmt.Println(newArr)

	arr = []int{0, 1, 2, 3, 4, 5}
	newArr = removeElement(arr, 5)
	fmt.Println(newArr)

	arr = []int{0, 1, 2, 3, 4, 5}
	newArr = removeElement(arr, 3)
	fmt.Println(newArr)
}
