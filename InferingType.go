package main

import "fmt"

// in  Go we didnot need to mention variable type, go infer it self from the values given, for float default is float64
// for int deafault is int32 or int64
func main() {
	// Defining daysOnVacation using := below:
	daysOnVacation := 4 // here we didn't needvto mention var

	// Defining hoursInDay using var and = below:
	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")
}
