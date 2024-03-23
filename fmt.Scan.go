package main

import "fmt"

func main() {
	fmt.Println("What would you like for lunch?")

	// Declare a var to store user input value
	var food string
	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)
}
