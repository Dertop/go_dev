package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   int
}

func (e Employee) Work() {
	fmt.Printf("Employee Name: %s, ID: %d is working.\n", e.Name, e.ID)
}

type Manager struct {
	Employee
	Department string
}

func main() {
	manager := Manager{
		Employee: Employee{
			Name: "Anuar",
			ID:   101,
		},
		Department: "Sales",
	}

	manager.Work()
}

/*
1. Embedding: Allows one struct to include another, enabling reuse of fields and methods.
Composition: Embedding supports composition by combining types to create more complex structures.

2. Method Promotion: Methods of the embedded struct can be called directly on the outer struct.
3. No: The outer struct's method will take precedence. The embedded type cannot override methods from the outer struct.
*/
