package main

import "fmt"

type StructVal struct {
	IntVal int
}

// Add passed by value, Struct would allocate new memory
func (s StructVal) Add() {
	fmt.Println("Func Add")
	fmt.Printf("Memory Location %p\n", &s)
	fmt.Printf("Value Before %+v\n", s)
	s.IntVal++
	fmt.Printf("Value After %+v\n", s)
}

// AddPbr passed by reference, would use the same loc and might change the original values
func (s *StructVal) AddPbr() {
	fmt.Println("Func Add Pass By Reference")
	fmt.Printf("Memory Location %p\n", s)
	fmt.Printf("Value Before %+v\n", s)
	s.IntVal++
	fmt.Printf("Value After %+v\n", s)
}

func main() {
	init := StructVal{
		IntVal: 10,
	}

	fmt.Printf("\nMain\nMemory Location:%p\nValue:%+v\n\n", &init, init)

	init.Add()
	fmt.Printf("\nAfter Function Add\nMemory Location:%p\nValue:%+v\n\n", &init, init)

	init.AddPbr()
	fmt.Printf("\nAfter Function AddPbr\nMemory Location:%p\nValue:%+v\n\n", &init, init)
}
