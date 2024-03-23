package main

import "fmt"

func main() {
	floatExample := 1.75
	// %T prints type of the variable value
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// %d prints float value
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years..", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// %.2f prints float value with 2 digits only after decimal
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
}
