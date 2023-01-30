package main

import "fmt"

func main() {

	// here is about declaring a slice rather than an array
	// for array declaring -> a := [3]int{1,2,3}
	a := []int{1, 2, 3}
	foo(a)
	fmt.Println(a)
}

func foo(sli []int) {
	sli[0] = sli[0] + 1
}
