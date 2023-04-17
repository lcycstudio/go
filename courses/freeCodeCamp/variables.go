/*
Agenda - Variables
- Variable declaration
- Redeclaration and shadowing
- Visibility
- Naming conventions
- Type conventions
*/

package main

import (
	"fmt"
	"strconv"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa((i))
	fmt.Printf("%v, %T\n", k, k)
}

/* Summary
- Naming conventions
  - Pascal or camelCase
  - Capitalize acronyms (HTTP, URL)
  - As short as reasonable
  - Longer names for longer lives
- Type conversions
  - destinationType(variable)
  - use strconv package for strings
*/
