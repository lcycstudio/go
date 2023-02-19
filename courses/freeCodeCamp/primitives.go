/*
Agenda
- Boolean type
- Numeric types
  - Integers
  - Floating point
  - Complex numbers
- Text types
*/

package main

import (
	"fmt"
)

func main() {
	var bo bool = true
	fmt.Printf("%v, %T\n", bo, bo)

	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	// go default c to false
	var c bool
	fmt.Printf("%v, %T\n", c, c)

	d := 42
	fmt.Printf("%v, %T\n", d, d)

	var e uint16 = 42
	fmt.Printf("%v, %T\n", e, e)

	// int
	f := 10
	g := 3
	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)
	fmt.Println(f % g)

	var h int = 10
	var i int8 = 3
	fmt.Println(h + int(i))

	j := 10             // 1010
	k := 3              // 0011
	fmt.Println(j & k)  // 0010
	fmt.Println(j | k)  // 1011
	fmt.Println(j ^ k)  // 1001
	fmt.Println(j &^ k) // 0100

	l := 8              // 2^3
	fmt.Println(l << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(l >> 3) // 2^3 / 2^3 = 2^0

	// float
	// n := 3.14
	var n float64 = 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	p := 10.2
	q := 3.7
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)

	// complex numbers
	var r complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", r, r)
	r = 1 + 2i
	s := 2 + 5.2i
	fmt.Println(r + s)
	fmt.Println(r - s)
	fmt.Println(r * s)
	fmt.Println(r / s)

	var t complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(t), real(t))
	fmt.Printf("%v, %T\n", imag(t), imag(t))

	var u complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", u, u)

	// strings
	v := "this is a string"
	fmt.Printf("%v, %T\n", string(v[2]), v[2])

	w := "this is also a string"
	fmt.Printf("%v, %T\n", v+w, v+w)

	// string to bytes
	x := []byte(v)
	fmt.Printf("%v, %T\n", x, x)

	// runes
	var y rune = 'a'
	fmt.Printf("%v, %T\n", y, y)
}

/* Summary
- Boolean type
  - Values are true or false
  - Not an alias for other types (e.g. int)
  - Zero value is false
- Nuermic types
  - Integers
	- Signed integers
	  - int type has varying size, but min 32 bits
	  - 8 bit (int8) through 64, bit (int64)
	- Unsigned integers
	  - 8 bit (byte and unit8) through 32 bit (unit32)
	- Arithmetic operations
	  - Addition, subtraction, multiplication, division, remainder
	- Bitwise operations
	  - And, or, xor, and not
	- Zero value is 0
	- Can't mix types in same family (unit16 + unit32 = error)
  - Floating point numbers
    - Follow IEEE-754 standard
	- Zero value is 0
	- 32 and 64 bit versions
	- Literal styles
	  - Decimal (3.14)
	  - Exponential (13e18 or 2E10)
	  - Mixed (13.7e12)
	- Arithmetic operations
	  - Addition, subtraction, multiplication, division
  - Complex numbers
    - Zero value is (0+0i)
	- 64 and 128 bit versions
	- Built-in functions
	  - complex - make complex number from two floats
	  - real - get real part as float
	  - imag - get imaginary part as float
	- Arithmetic operations
	  - Addition, subtraction, multiplication, division
- Text Types
  - Strings
    - UTF-8
	- Immutable
	- Can be concatenated with plus (+) operator
	- Can be converted to []byte
  - Rune
    - UTF-32
	- Alias for int32
	- Special methods normally required to process
	  - e.g. strings.Reader#ReadRune
*/
