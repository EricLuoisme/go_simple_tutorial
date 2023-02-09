package main

import (
	"fmt"
	"sync"
)

/*
Explanation:

Here is an example about race-condition that might happen and causing problems.
Whole program is about summing up two number and get the result. But instead of
adding the whole number, the number is adding 1 each times, then reach the result
e.g. 10000 -> for loop 10000 times and add 1 each loop. Two number inputs are set
15000 in the program. Now, Two goroutines is created for handling this summing task
separately, and verify any race-condition happens. Main procedure was loop 10 times
for better observation.

As you could see after running the program, the result could hardly get 30,000,
(In my computer, I never reach it), most the result of summation would between
[15,000, 30,000).

The reason why this happens, is because the two goroutines are in race-condition.
The Function addingOneByOne, which the goroutines calling, manipulate same global
variable 'summation'. In computer, program runs by COPYING data from the memory
to the CPU cache, then letting the CPU execute it, and finally copying the result
back to the memory for the program to use. In multi-core CPU, each core would execute
task independently. So in this case, two goroutines are considered as different tasks.
For a calling to Function addingOneByOne, OS would help COPY the value 'summation'.
If these two goroutines just 'run at the same time', then OS would COPY value TWICE
into two different CPU core, then inside this two cpu, val would +1 for each. After
summation, OS would COPY both value back to same slot, which means it +1 +1, but only
resulting in +1. That's the race-condition and why we only get result that less than
correct answer of the summation.

*/

var summation int

func main() {
	for i := 0; i < 10; i++ {
		summation = 0
		racingTest()
		fmt.Printf("For the %v test, summation got %v\n", i, summation)
	}
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
