package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	xi := []int{2, 4, 5, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 432,
		"Job":  34,
	}
	fmt.Println(m)

	jone := person{
		fname: "Steven",
		lname: "Job",
	}
	fmt.Println(jone)
	jone.speak()

	sa1 := secretAgent{
		jone,
		false,
	}
	sa1.speak()

	saySomething(jone)
	saySomething(sa1)
}

func (p person) speak() {
	fmt.Println(p.fname, `says : Good Morning`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, `says: Good Morning`)
	fmt.Println(sa.licenseToKill)
}
