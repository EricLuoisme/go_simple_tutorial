package main

import (
	"fmt"
	"strconv"
)

// LocalURL can be used from other pkg
var LocalURL = "http://127.0.0.1"

func main() {

	// must use two-steps init
	fmt.Println()
	fmt.Printf("%v, %T", LocalURL, LocalURL)

	// two-steps init
	var i int
	i = 42
	fmt.Println(i)

	// one-step init
	j := 3
	fmt.Printf("%v, %T", j, j)

	// type convert, must be converted by programmer
	var k float32
	k = float32(i)
	fmt.Println()
	fmt.Printf("%v, %T", k, k)

	// convert int into string, not to ASK-II must use strconv
	str := strconv.Itoa(i)
	fmt.Println()
	fmt.Printf("%v, %T", str, str)
}
