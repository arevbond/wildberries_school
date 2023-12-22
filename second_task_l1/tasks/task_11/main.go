/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	numToCount1 := make(map[int]int)
	for _, n := range nums1 {
		numToCount1[n]++
	}

	result := make([]int, 0)
	for _, n := range nums2 {
		if _, ok := numToCount1[n]; ok {
			result = append(result, n)
			numToCount1[n]--
			if numToCount1[n] == 0 {
				delete(numToCount1, n)
			}
		}
	}
	return result
}

func main() {
	nums1, nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{9, 4, 5, 3}
	fmt.Println(intersect(nums1, nums2))
}
