package main

import "fmt"

type Car struct {
	make  string
	model string
	year  int
}

func main() {
	// You can experiment with the Car struct if you wish.
	tata := Car{"SUV", "safari", 2020}
	fmt.Println(tata.model)
}
