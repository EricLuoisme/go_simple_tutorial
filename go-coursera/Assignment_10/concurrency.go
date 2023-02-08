package main

import (
	"fmt"
	"sync"
)

/*
Explanation:

Here is an example about race-condition that might happen and causing problems.
Whole program is about switching a string variable to 'A' 10,000 times and switch
it to 'B' another 10,000 times. sync.WaitGroup is used for let the function wait
till other tasks are finished. According to the order of the code, it should just
set to 'A' first (in a period), then result in variable == 'B'. Now, Two goroutines
is created for handling this two tasks separately.

But as you can see in the main function, we find out variable resulting in 'A' in some
case (in my computer I go 8/10 times of 'A'), which is different from the result that
only execute task in single goroutine

This is because, two goroutines manipulating the same variable, we

*/

var summation int

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

func racingTest() {
	// define waiting group for holding main function
	var wg sync.WaitGroup
	// two goroutines
	wg.Add(2)
	go addingOneByOne(&wg, 15000)
	go addingOneByOne(&wg, 15000)
	// waiting until all goroutines finished
	wg.Wait()
}

func addingOneByOne(wg *sync.WaitGroup, addVal int) {
	for i := 0; i < addVal; i++ {
		summation++
	}
	wg.Done()
}
