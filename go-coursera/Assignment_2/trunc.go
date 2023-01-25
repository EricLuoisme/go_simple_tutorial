package main

import "fmt"

func main() {
	var floatInput float32
	fmt.Println("Please enter a floating point number")

	_, err := fmt.Scan(&floatInput)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(int(floatInput))
}
