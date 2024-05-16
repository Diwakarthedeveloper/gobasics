package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	fmt.Println(donuts)

	// delete key:value pair using this format -> delete(MapName, KeyName)
	delete(donuts, "chocolate")
	fmt.Println(donuts) // printing updated Map

}
