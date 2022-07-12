package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	msg := "Go Routine"
	// tell main thread to wait
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// consume the waiting
		wg.Done()
	}(msg)
	msg = "Go Routine 2"
	// now wait
	wg.Wait()
}
