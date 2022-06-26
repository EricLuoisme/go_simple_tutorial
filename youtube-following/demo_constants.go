package main

import "fmt"

// iota means auto-increment, start from 0
const (
	// means do not care about 0
	_ = iota + 5
	b // 6
	c // 7
)

const (
	error = iota
	pending
	success
	failure
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmeria
	canSeeSouthAmeria
)

func main() {

	// for constant, it has to be init by settled type
	// it can not be changed, must could be determined before compile-time
	// it can also 'be shadowed'
	const myConst int = 42
	fmt.Printf("Const: %v, %T\n", myConst+23, myConst)

	fmt.Printf("b:%v, c:%v\n", b, c)

	// by using iota to act as const, need to make sure 'iota' start with error val
	// or else, some un-init variable would accidentally 'equal' to it
	var enumTest int
	fmt.Printf("Enum test with init: %v\n", enumTest == pending)

	// this means, we can make a complex 'switch' in a single byte data space
	// below roles, means it's admin + can see financials + can see europe
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("Roles: %b\n", roles)
	fmt.Printf("Can see south Africa ?: %v\n", canSeeAfrica&roles == canSeeAfrica)
}
