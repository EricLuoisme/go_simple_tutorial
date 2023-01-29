package main

import "fmt"

func main() {

	arrInt := []int{1, 2, 3, 4, 5}
	sliceInt := arrInt[3:]

	fmt.Println("Slice Example")
	fmt.Printf("Init\nArraInt: %+v, SliceInt: %+v\n\n", arrInt, sliceInt)
	sliceInt[0] = 10
	fmt.Printf("After\nArraInt: %+v, SliceInt: %+v\n\n", arrInt, sliceInt)
	sliceInt = append(sliceInt, 10)
	fmt.Printf("Final\nArraInt: %+v, SliceInt: %+v\n\n", arrInt, sliceInt)

	emptyMap := make(map[string]interface{})
	fmt.Println("Map Example")
	fmt.Printf("Init\nemptyMap : %+v\n", emptyMap)

	MapFunc(emptyMap)
	fmt.Printf("After\nemptyMap : %+v\n", emptyMap)
	// copy() can be used to copy obj in Maps work as 'Deep Copy'
}

func MapFunc(val map[string]interface{}) {
	val["this is a new value pair"] = 100
}
