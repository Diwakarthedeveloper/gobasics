package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out all the donuts
	fmt.Println(donuts)

	// AAdding new key:value pair to map in this format --> mapname["NewKey"] = NewValue
	donuts["glazed"] = 12
	fmt.Println(donuts["glazed"])

	//Modifying existing map in this format -->  mapname["ExistingKey"] = NewValue
	donuts["jelly"] = 3
	fmt.Println(donuts)

}
