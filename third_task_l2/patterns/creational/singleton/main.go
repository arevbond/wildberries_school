package main

import (
	"sync"
)

type President struct {
}

var (
	instance *President
	// Once is an object that will perform exactly one action.
	//
	// A Once must not be copied after first use.
	//
	// In the terminology of the Go memory model,
	// the return from f “synchronizes before”
	// the return from any call of once.Do(f).
	once sync.Once
)

func GetInstance() *President {
	once.Do(func() {
		instance = &President{}
	})
	return instance
}
