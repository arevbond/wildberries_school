/*
Реализовать конкурентную запись данных в map
*/

package main

import (
	"errors"
	"fmt"
	"sync"
)

var numbers = make(map[int]int, 100)

type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

func (s *SafeNumbers) Add(num int) {
	s.Lock()
	defer s.Unlock()
	s.numbers[num] = num
}

func (s *SafeNumbers) Get(num int) (int, error) {
	s.RLock()
	defer s.RUnlock()
	if number, ok := s.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("numbers doesn't exists")
}

func main() {
	generateNumbersMap()
}

func generateNumbersMap() {
	wg := sync.WaitGroup{}

	safenumbers := &SafeNumbers{
		numbers: numbers,
	}
	// Write.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safenumbers.Add(i)
		}(i)
	}
	// Read.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num, err := safenumbers.Get(i)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%d ", num)
				if num > 0 && num%20 == 0 {
					fmt.Println()
				}
			}
		}(i)
	}
}
