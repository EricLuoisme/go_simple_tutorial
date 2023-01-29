package main

import "fmt"

func main() {
	arr := []string{"hi", "bye"}

	fmt.Printf("arr: %v\n", arr)
	arrValFunc(arr)
	fmt.Printf("arr: %v\n", arr)
	arrRefFunc(&arr)
	fmt.Printf("arr: %v\n", arr)
}

// arrValFunc passing the array, would only change the value in the slot
// (because the array slot store the value's reference), but would not change
// the array, i.e. append would not affect
func arrValFunc(c []string) {
	for index, _ := range c {
		c[index] = "hello"
	}
	c = append(c, "seeU")
}

// arrRefFunc passing the pointer of an array, value's reference could be changed,
// also the array could be changed
func arrRefFunc(c *[]string) {
	for index, _ := range *c {
		(*c)[index] = "welcome"
	}
	*c = append(*c, "seeU")
}
