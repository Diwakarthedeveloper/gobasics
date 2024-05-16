package main

import "fmt"

// Country struct below with multiple variables related to country and their types
type Country struct {
	name      string
	capital   string
	latitude  float32
	longitude float32
}

func main() {
	var france Country
	fmt.Println(france)
}
