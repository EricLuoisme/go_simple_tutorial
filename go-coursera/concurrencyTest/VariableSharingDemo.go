package main

import (
	"fmt"
	"sync"
)

var i = 0
var wg sync.WaitGroup

func inc() {
	i = i + 1
	wg.Done()
}

func main() {
	wg.Add(2)
	// interleaving inside
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
