/*
Agenda
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

// uppercase is going to export
// lowercase is going to import
// use Pascal casing and camel casing
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
	Episodes   []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Lewis Chen",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.Number)
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.Companions)
	fmt.Println(aDoctor.Companions[0])
	fmt.Println(aDoctor.Episodes)

	bDoctor := struct{ name string }{name: "Bat Man"}
	cDoctor := bDoctor
	cDoctor.name = "Tom Baker"
	fmt.Println(bDoctor)
	fmt.Println(cDoctor)
}

/* Summary
-
*/
