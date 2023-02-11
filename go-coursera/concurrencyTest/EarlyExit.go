package main

import (
	"fmt"
)

func main() {
	// almost main finished before the new goroutine started
	go fmt.Println("New routine")
	fmt.Println("Main routine")
}
