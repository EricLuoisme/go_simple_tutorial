package main

import "fmt"

func main() {

	// in Golang, interfaces no using by implementation
	// one can 'implement' a interface, without actually 'implement'
	// all it's methods
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World!"))

}

// Writer interface describe behaviours
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter is a struct that have an 'behavior', which is a method,
// that 'implement' the Writer interface's Write Method
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
