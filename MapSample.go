package main

import "fmt"

func main() {
	// Below is a simple contact list.
	contacts := map[string]int{ // variable := map[key type]{value type}
		"Joe":    2126778723,
		"Angela": 4089978763,
		"Shawn":  3143776876,
		"Terell": 5026754531,
	}
	// Print out all the contacts
	fmt.Println(contacts)
}
