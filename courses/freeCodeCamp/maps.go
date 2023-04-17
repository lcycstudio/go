/*
Agenda - Maps and Structs
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

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	statePopulations["Georgia"] = 10310371
	// delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// check if a key exists
	_, ok := statePopulations["Oho"]
	fmt.Println(ok)

	// delete "Ohio" from both sp and statePopulations
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}

/* Summary
- Maps
  - Collections of value types that are accessed via keys
  - Created via literals or via make function
  - Members accessed via [key] syntax
	- myMap["key"] = "value"
  - Check for presence with "value, ok" form of result
  - Multiple assignments refer to same underlying data
*/
