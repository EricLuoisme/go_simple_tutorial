package main

import (
	"fmt"
	"strconv"
)

func main() {

	// go would initialize all variables
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	m := 42
	fmt.Printf("%v, %T\n", m, m)

	// un-sign integer
	var u uint8 = 42
	fmt.Printf("%v, %T\n", u, u)

	a := 10
	b := 3
	fmt.Printf("Ops variables: %v, %v\n", a, b)
	fmt.Println("And Ops: " + strconv.Itoa(a&b))
	fmt.Println("Or Ops: " + strconv.Itoa(a|b))
	fmt.Println("异或 Ops: " + strconv.Itoa(a^b))

	// not changing the original val in the address
	fmt.Println(a >> 1)
	fmt.Println(a)

	// complex type
	var com complex64 = 10 + 2i
	fmt.Printf("Complex Num: real:%v, imag:%v, type:%T\n", real(com), imag(com), com)

	// for string, go still store it as ASC-II
	// we could check it by converting it into bytes-slice
	// for different programming language communicated with go, might better using string -> byte to transfer info
	s := "This is a string"
	by := []byte(s)
	fmt.Printf("String test, number[2] in slot 3: %v, %v, %v\n", s[2], string(s[2]))
	fmt.Printf("Byte for string, %v\n", by)

	// rune type, can read some stuff which encoded into utf-32
	var r rune = 'a'
	fmt.Printf("Rune type, %v, %T\n", r, r)
}
