package main

import "fmt"

func main() {
	// defer call evaluated immediately
	i := 1
	defer fmt.Println(i + 1)
	i++
	i++
	i++
	fmt.Println("Hello?")
}
