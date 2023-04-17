/*
Agenda - Pointers
- Creating pointers
- Dereferencing pointers
- The new function
- Working with nil
- Types with internal pointers
*/

package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	// output: 42 0xc0000140c8
	fmt.Println(&a, b)
	// output: 0xc0000140c8 0xc0000140c8
	fmt.Println(a, *b)
	// output: 42 42
	a = 27
	fmt.Println(a, b)
	// output: 27 0xc0000140c8
	fmt.Println(a, *b)
	// output: 27 27
	*b = 14
	fmt.Println(a, *b)
	// output: 14 14

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)
	// output: [1 2 3] 0xc000018228 0xc000018230

	var ms *myStruct
	fmt.Println(ms)
	// output: <nil>
	// ms = &myStruct{foo: 42}
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println((*ms).foo)
	// output: 42
	ms.foo = 42
	fmt.Println(ms.foo)
	// output: 42

	g := [3]int{1, 2, 3}
	h := g
	fmt.Println(g, h)
	// output: [1 2 3] [1 2 3]
	g[1] = 42
	fmt.Println(g, h)
	// output: [1 42 3] [1 2 3]

	i := map[string]string{"foo": "bar", "baz": "buz"}
	j := i
	fmt.Println(i, j)
	// output: map[baz:buz foo:bar] map[baz:buz foo:bar]
	i["foo"] = "qux"
	fmt.Println(i, j)
	// output: map[baz:buz foo:qux] map[baz:buz foo:qux]

}

type myStruct struct {
	foo int
}

/* Summary
- Creating pointers
  - Pointer types use an asterisk (*) as a prefix to type pointed to
    - *int - a pointer to an integer
  - Use the addressof operator (&) to get address of variable
- Dereferencing pointers
  - Dereference a pointer by preceding with an asterisk (*)
  - Complex types (e.g. structs) are automatically dereferenced
- Create pointers to objects
  - Can use the addressof operator (&) if value type already exists
    - ms := myStruct{foo: 42}
	- p := &ms
  - Use addressof operator before initializer
    - &myStruct{foo: 42}
  - Use the new keyword
	- Can't initialize fields at the same time
- Types with internal pointers
  - All assignment operations in Go are copy operations
  - Slices and maps contain internal pointers, so copies point to same underlying data
*/
