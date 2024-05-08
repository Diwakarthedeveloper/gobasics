package main

import (
	"fmt"
)

func main() {

	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")
	// LOOP GOING THROUGH MENU BELOW THIS LINE USING range KEYWORD
	for number, item := range menu {
		fmt.Println("Number:", number, "Menu:", item)
	}

	orders := map[string]string{ //maps is just like indexed araay of strings in a key value pair but here is no indexing
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}
	// A late order
	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")
	// WRITE LOOP GOING THROUGH ORDERS BELOW THIS LINE
	for person, food := range orders {
		fmt.Println("Person", person, "Food", food)
	}

}
