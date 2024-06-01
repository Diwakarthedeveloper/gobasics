package main

import (
	"fmt"
)

func main() {

	for count := 1; count <= 12; count = count + 2 {
		fmt.Println(count)
	}

	// for loop with range
	nums := []int{1, 2, 3, 4, 5}
	for a, b := range nums { // here values is stored in the bums variable and used with the range function
		fmt.Printf("Index: %d, Value: %d\n", a, b)

	}
	fmt.Println("for Loop with range (for strings):")
	for a, b := range "INDIA" { // here values is not stored in the variable but directly used with the range function
		fmt.Printf("Index: %d, Rune: %c\n", a, b)
	}

}
