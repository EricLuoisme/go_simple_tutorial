package main

import (
	"fmt"
	"math"
)

func main() {

	// simple if statement
	if true {
		fmt.Println("The test is true")
	}

	// if statement with extra code for getting result
	stateProperties := map[string]int{
		"California": 3124,
		"Texas":      234,
		"Florida":    23432,
	}
	if pop, ok := stateProperties["Florida"]; !ok {
		fmt.Println(pop)
		fmt.Println(ok)
	}
	if returnTrue() {
		fmt.Println(getPower())
	}

	getSwitch()

	typeSwitch()

}

func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}

func getPower() float64 {
	return math.Pow(math.Sqrt(4), 2)
}

func getSwitch() {
	switch i := 1 + 3; i {
	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4, 6:
		fmt.Println("even")
	default:
		fmt.Println("no")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("<= 10")
		// keyword help to go to next case,
		// never mind it pass this current case or not
		fallthrough
	case i <= 20:
		fmt.Println("<= 20")
	default:
		fmt.Println("> 20")
	}
}

func typeSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is int")
	case float64:
		fmt.Println("i is float")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
