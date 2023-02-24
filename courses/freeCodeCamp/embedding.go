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

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // inheriting from the Animal structure
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	c := Bird{
		Animal:   Animal{Name: "Ume", Origin: "New Zealand"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(c.Name)
}

/* Summary
-
*/
