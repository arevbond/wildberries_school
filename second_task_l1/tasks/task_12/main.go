/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

package main

import "fmt"

func main() {
	strs := []string{"cat", "cat", "dog", "tree"}
	set := make(map[string]struct{})
	for _, str := range strs {
		if _, ok := set[str]; !ok {
			set[str] = struct{}{}
		}
	}
	fmt.Println(set)
}
