package main

import "fmt"

func main() {

	// simple array with initialization
	grades := [...]int{12, 4, 59}
	fmt.Printf("Grades: %v", grades)
	fmt.Printf("\n[1]: %v\n", grades[0])

	// init array without only length
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Tom"
	fmt.Printf("Students: %v, length: %v\n", students, len(students))

	// 2d arrays
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// array copy = deep copy (with ... size)
	// if we got too many args in array, copy would be really slow
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	c := &a // using pointer
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// but using slice, it would be shallow copy
	// because slice is naturally a referencing type
	d := []int{1, 2, 3}
	e := d
	e[1] = 9
	fmt.Println(d)
	fmt.Println(e)

	// slice real using
	fmt.Println("\nSlice test")
	aReal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bReal := aReal[:]   // slice of all elements
	cReal := aReal[3:]  // from the 4th element till the end
	dReal := aReal[:6]  // slice first 6 elements
	eReal := aReal[3:6] // 4, 5, 6
	aReal[5] = 1000
	// all below would contain 1000
	fmt.Println(aReal)
	fmt.Println(bReal)
	fmt.Println(cReal)
	fmt.Println(dReal)
	fmt.Println(eReal)

	// make func
	fmt.Println("\nMake Function")
	aMake := make([]int, 3, 6)
	fmt.Printf("AMake:%v, Length:%v, Capacity:%v\n", aMake, len(aMake), cap(aMake))
	// append and increase the length, but capacity still can hold, thus, capacity remain the same
	bMake := append(aMake, 1, 2, 3)
	fmt.Printf("BMake:%v, Length:%v, Capacity:%v\n", bMake, len(bMake), cap(bMake))
	// concat slices, ... means add them one by one
	cMake := append(bMake, aMake...)
	fmt.Printf("CMake:%v, Length:%v, Capacity:%v\n", cMake, len(cMake), cap(cMake))

	// need to take care of slice removing
	fmt.Println("\nSlice Removing")
	dMake := []int{1, 2, 3, 4, 5}
	fmt.Println(dMake)
	eMake := append(dMake[:2], dMake[3:]...)
	fmt.Println(eMake)
	// could find slice's from a slice, if they change the val
	// because of shallow copy, it would also 'pollute' the original slice
	fmt.Println(dMake)

}
