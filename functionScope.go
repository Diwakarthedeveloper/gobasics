package main

import "fmt"

//to show a variable is only assiccible within a function not outside  a function unless function type and return staetement is given

func startGame() {
	instructions := "Press enter to start..."
	fmt.Println(instructions) // instructions is only accisible inside function startGame

}

func main() {
	startGame()
	//fmt.Println(instructions)  // this will throw error here as instructions is defined in startGame function and without return type only vales can be called only insdie a function.

}
