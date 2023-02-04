package main

import (
	"fmt"
	"math"
)

func main() {

	// 1. prepare values for acceleration, initial velocity, and displacement
	var a, v0, s0 float64
	if !readInput("acceleration value a", &a) ||
		!readInput("initial velocity v_0", &v0) ||
		!readInput("initial displacement s_0", &s0) {
		return
	}

	fmt.Printf("%v, %v, %v\n", a, v0, s0)

	// 2. calculate formula
	formula := GenDisplaceFn(a, v0, s0)

	// 3. input & calculate formula output
	var t float64
	if !readInput("time t", &t) {
		return
	}
	fmt.Println(formula(t))
}

// GenDisplaceFn would return a formula -> (1/2)at^2 + v_0t + s_0
func GenDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

// readInput would prompt string & read float64 input
func readInput(notice string, pointer *float64) bool {
	fmt.Printf("Please enter value for %s\n", notice)
	_, err := fmt.Scan(pointer)
	if err != nil {
		fmt.Println("Error occurred, Stop the program")
		return false
	}
	return true
}
