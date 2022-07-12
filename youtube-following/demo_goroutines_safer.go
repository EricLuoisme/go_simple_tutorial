package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Go Routine"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	// after sending a value (hard copy in golang)
	// we can make sure thread2 would definitely print out what we input before
	msg = "Go Routine 2"
	time.Sleep(100 * time.Millisecond)
}
