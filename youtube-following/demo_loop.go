package main

import "fmt"

func main() {
	// simple loop
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// multi-val loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("%v, %v  ", i, j)
	}

	// traverse slice without knowing its siz
	s := []int{1, 2, 3}
	for key, val := range s {
		fmt.Printf("Key %v, Val %v\n", key, val)
	}

}
