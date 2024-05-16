package main

import "fmt"

func main() {
	// Creating a map with make
	orders := make(map[string]float32)
	fmt.Println(orders)
	//Creating a map with values
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 125,
		"jelly":     8,
	}
	fmt.Println(donuts)
}
