/*
Agenda - Looping
- For statements
  - simple loops
  - exiting early
  - looping through collections
*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
		fmt.Println(i)
	}

	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s[1])
	for k, v := range s {
		fmt.Println(k, v)
	}

	t := "hello Go!"
	for k, v := range t {
		fmt.Println(k, string(v))
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}
}

/* Summary
- For statements
  - simple loops
    - for initializer; test; incrementer {}
    - for test {}
    - for {}
  - exiting early
    - break
    - continue
    - labels
  - looping over collections
    - arrays, slices, maps, strings, channels
    - for k, v := range collection {}
*/
