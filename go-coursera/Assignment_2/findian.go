package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// scan input, could read whole line by using Buffered-Scanner
	fmt.Println("Please enter")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()

	// case-insensitive
	inputStr = strings.ToLower(inputStr)

	// three conditions, might be faster if using hard coding with i && I, n && N, a && A
	result := "Not Found!"
	if strings.HasPrefix(inputStr, "i") &&
		strings.HasSuffix(inputStr, "n") &&
		strings.ContainsAny(inputStr, "a") {
		result = "Found!"
	}

	// print
	fmt.Println(result)
}
