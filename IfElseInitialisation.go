package main

import "fmt"

func main() {

	// in belwo code y is defined within if block so it will not be accessible from outside the y block

	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	}
	// Note: y is not accessible outside the if block
	//fmt.Println(y) // this will throw error and synayx error as y is not found globally in the program

}
