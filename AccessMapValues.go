package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out the inventory
	fmt.Println(donuts)

	// variable := yourMap[keyValue] <-- Format to accss map and store in variable
	firstChoice := donuts["frosted"]
	fmt.Println(firstChoice)

	//secondChoice, status := donuts["bavarian cream"]
	secondChoice, status := donuts["chocolate"]

	fmt.Println(status)
	if status {
		fmt.Println(secondChoice)
	} else {
		fmt.Println("We do not sell that donut!")
	}

}
