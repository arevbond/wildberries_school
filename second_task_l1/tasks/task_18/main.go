/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить
итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	counter := &Counter{value: 0}

	quit := make(chan bool)

	n := 33

	go Do(counter, n, quit)

	<-quit
	fmt.Printf("Итоговое значение счётчика при %d горутин: %d\n", n, counter.Value())
}

func Do(counter *Counter, n int, quit chan bool) {
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int, c *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d increment value\n", i+1)
			c.Inc()
		}(i, counter, &wg)
	}
	wg.Wait()
	quit <- true
	close(quit)
}
