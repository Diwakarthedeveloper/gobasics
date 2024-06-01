//we can also provide functions with information using parameters. Function parameters are variables that are used within the function to use in some sort of computation or calculation. When calling a function, we give arguments, which are the values that we want those parameter variables to take.

package main

import "fmt"

// Update marsYear so that it takes earthYears
// As a parameter
func computeMarsYears(earthYears int) int { //  last int is the data type of value to be returned

	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func main() {
	myAge := 25 // go takes myAge as earthYears when myage is passed below in function online 20

	// Calling `marsYear` with `myAge`
	myMartianAge := computeMarsYears(myAge)
	fmt.Println(myMartianAge)
}
