package main

import "fmt"

func main() {
	// channel can be reused
	c := make(chan int)
	// put more into channel would be all good
	go prod(1, 2, c)
	go prod(3, 4, c)
	// if we receive more than channel produced data, it would become deadlock
	a := <-c
	b := <-c
	fmt.Println(a * b)
}

func prod(v1, v2 int, c chan int) {
	c <- v1 * v2
}
