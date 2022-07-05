package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	// copy as new var
	a := 42
	b := a
	fmt.Println(a, b)
	a = 3
	fmt.Println(a, b)
	fmt.Println()

	// pointer copy
	var c int = 42
	var d *int = &c
	fmt.Println(c, d)
	// get address &
	fmt.Println(&c, d)
	fmt.Println(&c, &d)
	c = 3
	fmt.Println(c, d)
	// de-referencing *
	fmt.Println(c, *d)
	fmt.Println()

	// array
	e := []int{1, 2, 3}
	f := &e[0]
	// but we can not directly do the pointer calculation
	g := &e[2]
	fmt.Printf("%v, %p, %p\n", e, f, g)

	// struct
	var ms *myStruct
	// get 'nil', which means we need to check it and 'panic'
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
	// de-reference the pointer, get the construct and sign val
	(*ms).foo = 42
	fmt.Println(ms)
	fmt.Println((*ms).foo)
	// but in golang, it can be decoded by the compiler
	ms.foo = 99
	fmt.Println(ms)
	fmt.Println(ms.foo)
}
