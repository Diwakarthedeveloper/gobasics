package main

import (
	"fmt"
)

func ask() int {
	var input int
	fmt.Print("I am thinking of a number between 1-100: ")
	// Here we Get the input from the user to open he door
	fmt.Scan(&input)
	return input
}

func main() {
	var guess int

	// below is the indefenite loop logic, the program will ask the question infinite times
	// till you give the correct answer which is 56

	for guess != 56 {
		guess = ask()
	}
	if guess == 56 {
		fmt.Println("You are correct! You may go through to the treasure!")
	}
}
