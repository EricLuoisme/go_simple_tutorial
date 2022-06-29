package main

import "fmt"

func main() {
	// simple map
	stateProperties := map[string]int{
		"California": 213495,
		"Texas":      23134,
		"Florida":    2134,
	}
	fmt.Println(stateProperties)

	// complex map
	twoDimensionProperties := map[[2]int]string{
		[2]int{1, 3}: "13",
		[2]int{2, 4}: "24",
	}
	fmt.Println(twoDimensionProperties)

	// use make func to construct map
	makeStateProperties := make(map[string]int)
	makeStateProperties = map[string]int{
		"California": 213495,
		"Texas":      23134,
		"Florida":    2134,
	}
	fmt.Println(makeStateProperties)
	fmt.Println(makeStateProperties["California"])
	fmt.Println(makeStateProperties["Cat"])
	// the 'put' ops in golang's map, just the same as adding into an arr
	makeStateProperties["Cat"] = 16
	fmt.Println(makeStateProperties["Cat"])
	delete(makeStateProperties, "Cat")
	fmt.Println(makeStateProperties["Cat"])
	val, found := makeStateProperties["Cat"]
	fmt.Printf("%v, %v\n", val, found)
	// map also share physical space
	newMap := makeStateProperties
	newMap["California"] = 101
	fmt.Println(makeStateProperties)
}
