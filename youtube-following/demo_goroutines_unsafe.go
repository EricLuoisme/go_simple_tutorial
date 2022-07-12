package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Go Routine"
	go func() {
		fmt.Println(msg)
	}()
	// here might print go routine 2, because even calling make a new
	// routine, but the main thread just keep running, and let the msg
	// convert to new val
	msg = "Go Routine 2"
	time.Sleep(100 * time.Millisecond)
}
