package main

import (
	"fmt"
	"math"
)

func main() {
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))
}

// MakeDistOrigin returning the func (float64, float64) float64 as a return
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		// o_x & o_y are in the closure of fn, thus, even MakeDisOrigin
		// already executed, fn is not, fn can still access them
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}
