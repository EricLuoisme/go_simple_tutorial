package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// init by make
	theSlice := make([]int, 3)

	// traverse the slice
	curLen := 0
	for {
		// 1. scanning input
		var inputStr string
		fmt.Println("input new integer or 'X' for exit")
		_, err := fmt.Scan(&inputStr)
		if err != nil {
			return
		}
		// 2. exit the loop if match "X"
		if strings.Compare("X", inputStr) == 0 {
			break
		}
		// 3. adding into slice
		atoi, _ := strconv.Atoi(inputStr)
		if curLen >= 3 {
			theSlice = append(theSlice, atoi)
		} else {
			theSlice[curLen] = atoi
		}
		curLen++
	}

	// sorting -> under two lines both working
	sort.Ints(theSlice)
	//sort.Slice(theSlice, func(i, j int) bool { return theSlice[i] < theSlice[j] })
	fmt.Println(theSlice)
}
