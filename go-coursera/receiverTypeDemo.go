package main

import "fmt"

type MyInt struct {
	x float64
	y float64
}

func (mi MyInt) doubleFunc() int {
	return int(mi.x*2 + mi.y)
}

func main() {
	v := MyInt{
		x: 100.3,
		y: 20.3,
	}
	fmt.Println(v.doubleFunc())
}
