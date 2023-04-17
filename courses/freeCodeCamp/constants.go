/*
Agenda - Constants
- Arrays
  - Creation
  - Built-in functions
  - Working with arrays

- Slices
  - Creation
  - Built-in functions
  - Working with slices
*/

package main

import (
	"fmt"
)

const (
	a = iota
	b // = iota
	c // = iota
)

const (
	a2 = iota
)

const (
	zeroSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	// _ : one and only write only variable
	_ = iota + 5 // fixed offset a
	redCrayon
	yellowCrayon
	pinkCrayon
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 2^(10 * 1)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// math.Sin(1.57) (value of type float64) is not constant
	// const myConst2 float64 = math.Sin(1.57)
	// fmt.Printf("%v, %T\n", myConst2, myConst2)

	// invalid operation: a + b (mismatched types int and int16)
	// const a int = 42
	// var b int16 = 27
	// fmt.Printf("%v, %T\n", a+b, a+b)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%v\n", a2)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)

	fmt.Printf("%v\n", redCrayon)

	fileSize := 4000000000.
	fmt.Printf("%v, %T\n", KB, KB)
	fmt.Printf("%v, %T\n", MB, MB)
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}

/* Summary
- Immutable, but can be shadowed
- Replaced by the complier at compile time
  - Value must be calculable at compile time
- Named like variables
  - PascalCase for exported constants
  - camelCase for interal constants
- Typed constants work like immutable variables
  - Can interoperate only with same type
- Untyped constants work like literals
  - Can interoperate with similar types
- Enumerated constants
  - Special symbol iota allows related constants to be created easily
  - Iota starts at 0 in each const block and increments by one
  - Watch out of constant values that match zero values for variables
- Enumerated expressions
  - Operations that can be determined at compile time are allowed
  	- Arithmetic
  	- Bitwise operations
	- Bitshifting
*/
