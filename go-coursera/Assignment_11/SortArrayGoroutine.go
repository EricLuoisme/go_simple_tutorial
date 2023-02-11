package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// 1. reading input
	strSlice := readSplitInputLine()
	intSlice := make([]int, len(strSlice))
	for i, v := range strSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}

	// 2. split task
	cn := make(chan []int, 4)
	each := len(strSlice) / 4
	go splitSort(intSlice[:each], cn)
	go splitSort(intSlice[each:each*2], cn)
	go splitSort(intSlice[each*2:each*3], cn)
	go splitSort(intSlice[each*3:], cn)

	// 3. Get result
	fmt.Println(wholeSorted(4, cn))
}

// wholeSorted uses to retrieve sorted slice form channel and combine them together
func wholeSorted(numOfParts int, cn chan []int) []int {
	var storing = <-cn
	for i := 0; i < numOfParts-1; i++ {
		if nil == storing {
			storing = <-cn
			continue
		}
		storing = merge2Slice(storing, <-cn)
	}
	return storing
}

// splitSort is simple sorting and signify to then channel after finished
func splitSort(part []int, cn chan []int) {
	fmt.Printf("split slice for this task: %v\n", part)
	sort.Ints(part)
	cn <- part
}

// merge2Slice uses to merging 2 sorted int slice and resulting combination of them
func merge2Slice(s1, s2 []int) []int {
	p1 := len(s1)
	p2 := len(s2)
	idx := p1 + p2
	var whole = make([]int, idx)
	for p1 > 0 && p2 > 0 {
		s1Max := s1[p1-1]
		s2Max := s2[p2-1]
		if s1Max > s2Max {
			whole[idx-1] = s1Max
			p1--
		} else {
			whole[idx-1] = s2Max
			p2--
		}
		idx--
	}
	if p1 > 0 {
		copy(whole[:p1], s1[:p1])
	}
	if p2 > 0 {
		copy(whole[:p2], s2[:p2])
	}
	return whole
}

// readSplitInputLine reading whole input line and split them into string slice
func readSplitInputLine() []string {
	// scan
	fmt.Println("Please entry integer which split by space, enter for finished. e.g. 5 15 13 7 22 9 1 3 2 6 4 8")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()
	// return split str-integers
	return strings.Fields(inputLine)
}
