/*
Agenda - If & Switch Statements
- If statements
  - Operators
  - If - else and if - else if statements
- Switch statements
  - Simple cases
  - Cases with multiple tests
  - Falling through
  - Type switches
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
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
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 30

	// passing a function inside the if statement
	// short circuited
	// guess := -5
	// if guess < 1 || returnTrue() || guess > 100 {
	// 	fmt.Println("The guess must be between 1 and 100!")
	// }
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("Abs: These are the same")
	} else {
		fmt.Println("Abs: These are different")
	}

	// switch statement
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("other number")
	}

	// initialize with tags
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("other number")
	}

	// taggless
	j := 10
	switch {
	case j <= 10:
		fmt.Println("less than or equal to 10")
		// fallthrough (want to execute the next case)
	case j <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var k interface{} = [3]int{}
	switch k.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

/* Summary
- If statements
  - Initializer
  - Comparison operators
  - Logical operators
  - Short circuiting
  - If - else statements
  - If - else if statements
  - Equality and floats
- Switch statements
  - Switching on a tag
  - Cases with multiple tests
  - Initializers
  - Switches with no tag
  - Fallthrough
  - Type switches
  - Breaking out early
*/
