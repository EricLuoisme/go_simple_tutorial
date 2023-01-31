package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("\nPlease input integers as one-line string, " +
		"which separated with space character, up to 10 integers, e.g. \"3 1 9 4 5234 0 2 5 689\" ")

	// 1. read input & conditioning
	strSlice := readSplitInputLine()
	if len(strSlice) > 10 {
		fmt.Println("Number of input integers > 10")
		return
	}

	// 2. convert to int slice
	intSlice := make([]int, len(strSlice))
	for i, v := range strSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}

	// 3. BubbleSort
	BubbleSort(intSlice)
	fmt.Println(intSlice)
}

// readSplitInputLine reading whole input line and split them into string slice
func readSplitInputLine() []string {
	// scan
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	// return split integers
	return strings.Fields(inputLine)
}

// BubbleSort using bubble sort algorithm sorting the input array from least to greatest
func BubbleSort(intSlice []int) {

	sliceLen := len(intSlice)
	for outerIdx := 0; outerIdx < sliceLen-1; outerIdx++ {
		noMoreSwap := true
		for innerIdx := 0; innerIdx < sliceLen-1; innerIdx++ {
			if intSlice[innerIdx] > intSlice[innerIdx+1] {
				Swap(intSlice, innerIdx)
				noMoreSwap = false
			}
		}
		// for faster stop if no more swap occurred
		if noMoreSwap {
			break
		}
	}
}

// Swap would swap the content of intSlice in position i and position i+1
func Swap(intSlice []int, i int) {
	intSlice[i], intSlice[i+1] = intSlice[i+1], intSlice[i]
}
