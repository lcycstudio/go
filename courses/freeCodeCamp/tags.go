/*
Agenda - Tags
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
	"reflect"
)

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal   // inheriting from the Animal structure
	SpeedKPH float32
	CanFly   bool
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

/* Summary
-
*/
