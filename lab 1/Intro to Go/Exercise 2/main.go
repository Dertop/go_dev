package main

import (
	"fmt"
)

func main() {
	var age int = 22
	var gpa float64 = 3.1
	var name string = "Anuar"
	var isStudent bool = true

	randomNumber := 65.5
	university := "KBTU"
	hasScholarship := false

	fmt.Printf("age: %d, type: %T\n", age, age)
	fmt.Printf("gpa: %.1f, type: %T\n", gpa, gpa)
	fmt.Printf("name: %s, type: %T\n", name, name)
	fmt.Printf("isStudent: %t, type: %T\n", isStudent, isStudent)

	fmt.Printf("randomNumber: %.1f, type: %T\n", randomNumber, randomNumber)
	fmt.Printf("university: %s, type: %T\n", university, university)
	fmt.Printf("hasScholarship: %t, type: %T\n", hasScholarship, hasScholarship)
}

/*
1. var : Explicit type declaration, usable in both global and local scopes. Can be declared without initialization.
:= : Implicit type inference, usable only in local scopes (inside functions), and requires immediate initialization.
2. fmt.Printf("Type: %T\n", variable)
3. No. Go is statically typed, so the type of variable is fixed after declaration.
*/
