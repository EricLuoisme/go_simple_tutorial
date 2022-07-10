package main

import (
	"fmt"
)

func main() {
	// when passing simple type (not slice, map nor pointers)
	// it just passing a hard-copy into the func
	answer := "Yeah"
	sayMessage("Happy?", answer)
	// thus, the calling function's val would not be changed
	fmt.Println(answer)

	// e.g. slice & pointers, then it would be changed
	fmt.Println()
	question := "How old are u"
	askQuestion(&question)
	fmt.Println(question)

	// multiple values
	fmt.Println()
	sumP := sum(1, 2, 3, 4, 5)
	fmt.Println(sumP)
	fmt.Println(*sumP)

	// simple return declare
	fmt.Println()
	fmt.Println(simpleReturn())

	// error handling
	fmt.Println()
	_, e := divided(45.0, 0)
	fmt.Println(e)

	// anonymous func
	fmt.Println()
	withAnonymous()

	// method invocation, can calling the function that only we have a struct
	g := greeter{greeting: "Hello", name: "Jack"}
	g.greet()
	g.greetP()
	fmt.Println(g.greeting)
}

// if there are two same type input, in golang u can save one for concise
// thinking about passing a hard copy would be time-consuming
func sayMessage(msg1, msg2 string) {
	fmt.Println(msg1)
	fmt.Println(msg2)
	msg2 = "Nope"
	fmt.Println(msg2)
}

// if passing pointer, map or slice, no more hard copy, changes would be remained
// but thinking about passing a pointer is much more efficient
func askQuestion(question *string) {
	fmt.Println(*question)
	*question = "Who are u"
	fmt.Println(*question)
}

// multiple same types of input can be wrapped as using '...'
// also, golang can return a pointer of the result
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, value := range values {
		result += value
	}
	// let it return pointer
	// detail: usually not all language let you return a pointer of a value that using in current stack
	// because after finished execution, this function would have no idea of what stored inside the slot that
	// the return pointer pointing at, or worse, slot had been cleared
	// but in golang, if you are going to return a pointer, it would 'push' this value into the 'shared memory'
	// which save the value (it's memory) being recycled
	return &result
}

// by declaring a name, we do not need to init a val for storing the result
func simpleReturn() (result int) {
	result = 5
	return
}

// error handling
func divided(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divided by 0")
	}
	return a / b, nil
}

// func with anonymous func
func withAnonymous() {
	// real anonymous func, the last two buckets means that running this func at the same time
	func() {
		fmt.Println("Anonymous!")
	}() // u can put val into this buckets like passing into to the anonymous func
}

// method invocation
type greeter struct {
	greeting string
	name     string
}

// we can see the struct greeter as a value receiver
// which also 'strict' this func would only be calling by this kind of struct
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// if we need to change the val, we need to use pointer become the value receiver
func (g *greeter) greetP() {
	g.greeting = "Akumatata"
}
