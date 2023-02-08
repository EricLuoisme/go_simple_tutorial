package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	// define waiting group for holding main function
	var wg sync.WaitGroup

	// two goroutines
	wg.Add(2)
	go addOneByOne(&wg, 10000)
	go addOneByOne(&wg, 10000)

	// waiting until all goroutines finished
	wg.Wait()
	fmt.Println(count)
}

func addOneByOne(wg *sync.WaitGroup, addVal int) {
	for i := 0; i < addVal; i++ {
		count++
	}
	wg.Done()
}
