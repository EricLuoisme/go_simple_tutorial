package main

import (
	"fmt"
	"testing"
)

func main() {

	// simple defer test
	// defer would be FILO, similar to Stack
	// thus, we were usually using the 'defer' keyword to close the resource
	defer fmt.Println("1")
	defer fmt.Println("2")
	// defer would be executed, just before the func return its value
	fmt.Println("3")

	deferTest()
	deferWithVar()
	panicTest()
}

func deferTest() {
	// defer would be executed, just before the func return its value
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

func deferWithVar() {
	a := "a"
	defer fmt.Println(a)
	// the var already set when reaching the line defer
	// thus, only print out "a"
	a = "b"
}

func panicTest() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("Something bad happen")
	fmt.Println("end")
}
