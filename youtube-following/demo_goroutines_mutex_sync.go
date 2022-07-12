package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// by default goroutine number = cpu core number
	fmt.Println(runtime.GOMAXPROCS(-1))
	// setting 100 goroutine
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wgg.Add(2)
		// must be added the lock before calling the func
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wgg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	// must add unlock after func execution
	m.RUnlock()
	wgg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wgg.Done()
}
