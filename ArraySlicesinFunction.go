package main

import "fmt"

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	changeLastElement(myTutors, "Bobby")
}

// below we define the function to take in a slice parameter and a string.
func changeLastElement(teacher []string, newName string) { // teacher is a slice parameter & newName is a new stricg value
	// below  changing the last element of the slice to the new string value.
	// The last element will be at the index of the length minus one.
	teacher[len(teacher)-1] = newName
	fmt.Println(teacher)
}
