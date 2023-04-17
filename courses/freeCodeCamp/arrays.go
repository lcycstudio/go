/*
Agenda - Arrays and Slices
- Maps
  - What are they?
  - Creating
  - Manipulation

- Structs
  - What are they?
  - Creating
  - Naming conventions
  - Embedding
  - Tags
*/

package main

import (
	"fmt"
)

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity Matrix: %v\n", identityMatrix)

	aa := [...]int{1, 2, 3}
	bb := aa
	bb[1] = 5
	fmt.Println(aa)
	fmt.Println(bb)

	cc := &aa
	cc[1] = 6
	fmt.Println(aa)
	fmt.Println(cc)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice 4th, 5th and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	a[5] = 42
	fmt.Println(a)

	h := make([]int, 3)
	fmt.Println(h)
	fmt.Printf("Length:  %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	j := []int{}
	fmt.Println(j)
	fmt.Printf("Length:  %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))
	j = append(j, 1)
	fmt.Println(j)
	fmt.Printf("Length:  %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))
	j = append(j, 2, 3, 4, 5)
	fmt.Println(j)
	fmt.Printf("Length:  %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	// concatenate new slice to j
	j = append(j, []int{6, 7, 8, 9}...)
	fmt.Println(j)
	fmt.Printf("Length:  %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	k := []int{1, 2, 3, 4, 5}
	l := append(k[:2], k[3:]...)
	fmt.Println(l)
	fmt.Println(k)
}

/* Summary
- Arrays
  - Collection of items with same type
  - Fixed size
  - Declaration styles
	- a := [3]int{1, 2, 3}
	- a := [...]int{1, 2, 3}
	- var a [3]int
  - Access via zero-based index
	- a := [3]int{1, 3, 5} // a[1] == 3
  - len function returns size of array
  - copies refer to different underlying data
- Slices
  - Backed by array
  - Creation styles
	- Slice existing array or slice
	- Literl style
	- Via make function
	  - a := make([]int, 10) // create slice with capacity and length == 10
	  - a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
  - len function returns length of slice
  - cap function returns length of underlying array
  - append function to add elements to slice
	- May cause expensive copy operation if underlying array is too small
  - Copies refer to same underlying array
*/
