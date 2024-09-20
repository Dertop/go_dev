package main

import (
	"fmt"
)

// Function that adds two integers and returns their sum
func add(a int, b int) int {
	return a + b
}

// Function that swaps two strings and returns them in reverse order
func swap(x, y string) (string, string) {
	return y, x
}

// Function that returns both the quotient and remainder of two integers
func divide(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	// Test Add function
	sum := add(10, 5)
	fmt.Printf("Sum: %d\n", sum)

	// Test Swap function
	first, second := swap("Hello", "World")
	fmt.Printf("Swapped: %s, %s\n", first, second)

	// Test divide function
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}

/*
1. func swap(x, y string) (string, string) {
	return y, x
}
2. Act as pre-declared variables, making the code cleaner and allowing for a return statement without arguments
3. Blank Identifier (_)
*/
