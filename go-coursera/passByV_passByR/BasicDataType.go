package main

import "fmt"

// Basic types like int, int8, unit15, float32, string, bool, byte, Array, Structs
func main() {
	a, b := 0, 0

	fmt.Println("\n## Init")
	fmt.Printf("Memory Location a: %p, b: %p\n", &a, &b)
	fmt.Printf("Value a:%d, b:%d\n", a, b)

	Add(a)

	AddRef(&b)

	fmt.Println("\n## Final")
	fmt.Printf("Memory Location a: %p, b: %p\n", &a, &b)
	fmt.Printf("Value a:%d, b:%d\n", a, b)

}

func Add(x int) {
	fmt.Println("\n## 'Add' Function")
	fmt.Printf("Before Add, Memory Location: %p, Value: %d\n", &x, x)
	x++
	fmt.Printf("After Add, Memory Location: %p, Value: %d\n", &x, x)
}

func AddRef(x *int) {
	fmt.Println("\n## 'Add By Reference' Function")
	fmt.Printf("Before Add Ref, Memory Location: %p, Value: %d\n", x, *x)
	*x++ // * will call the value of a pointer
	fmt.Printf("After Add Ref, Memory Location: %p, Value: %d\n", x, *x)
}
