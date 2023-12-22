/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func writeInChannel(c chan string) {
	defer wg.Done()
	for {
		c <- fmt.Sprintf("some data %d", rand.Intn(1e3))
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func readFromChannel(indx int, c chan string) {
	for s := range c {
		fmt.Printf("%d worker read data: \"%s\"\n", indx, s)
	}
}

var wg = sync.WaitGroup{}
var nWorker int

func init() {
	flag.IntVar(&nWorker, "n", 10, "amount workers")
}

func main() {
	flag.Parse()
	c := make(chan string)

	wg.Add(1)
	go writeInChannel(c)

	for i := 1; i <= nWorker; i++ {
		go readFromChannel(i, c)
	}

	wg.Wait()
}
