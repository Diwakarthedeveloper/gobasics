package main

import "fmt"

func main() {

	const earthsGravity = 9.80665 // Name of Constant should be in camelCase or PascalCase
	//earthsGravity = 34 this will throw error as in above line its constant not variable, Constant cannot be changed once declared.

	fmt.Println("gravity of earth is ", earthsGravity)
}
