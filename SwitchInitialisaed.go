package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("switch with Initialization:")
	switch today := time.Now().Weekday(); today { // variable defined inside loop which is now not accisible outside
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("switch Without a Condition (like if-else):")
	x := 42
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is positive")
	}

}
