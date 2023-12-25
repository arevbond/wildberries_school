/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("sleep...")
	Sleep(3)
	fmt.Println("Wake up!")
}
