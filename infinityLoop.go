package main

import (
	"fmt"
)

var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)
	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {

	// below is the infinity loop containing the count function
	for {
		count()
	}

}
