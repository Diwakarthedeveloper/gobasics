package main

import "fmt"

func main() {
	fmt.Println("Below you see value of variable, address of Variable using pointer, and pointer exactrating value of variable. ")
	var a int = 314
	var x *int = &a

	fmt.Println("value of variable a = ", a)
	fmt.Println("Memory address of variable a", x)
	fmt.Println("Pointer exactrating value of variable = ", *x)
}

//output
// Below you see value of variable, address of Variable using pointer, and pointer exactrating value of variable.
// value of variable a =  314
// Memory address of variable a 0xc00000a0a8
// Pointer exactrating value of variable =  314
