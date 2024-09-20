package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{
		Name: "Anuar",
		Age:  22,
	}

	person.Greet()
}

/*
1. A struct in Go is defined using the type keyword followed by
the struct name and the struct keyword. Fields within the struct
are defined with their names and types.
2. Methods have a special receiver argument that allows
them to be associated with a specific type. This receiver
is specified between the func keyword and the method name.

Methods are associated with instances of a type, while
regular functions are not.

3. Yes. Methods can be associated with any user-defined type,
including types based on primitives, slices, or maps.
*/
