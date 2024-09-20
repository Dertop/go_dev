package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	var number int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d", &number)

	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	// Exercise 2
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("The sum of the first 10 natural numbers is: %d\n", sum)

	// Exercise 3
	var day int
	fmt.Print("Enter a number (1-7) to get the corresponding day of the week: ")
	fmt.Scanf("%d", &day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input! Please enter a number between 1 and 7.")
	}
}

/*
1. Optional initialization and no parentheses
2. Standard loop
While loop
Infinite loop
3.Multiple cases per Line and no need for break
*/
