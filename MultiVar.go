package main

import "fmt"

func main() {
	// Here we will learn to drfine multiple variable in one line
	var magicNum, powerLevel int32 // both var has same type but := is no used

	magicNum = 2048
	powerLevel = 9001
	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

	// Define amount and unit below:
	amount, unit := 10, "doll hairs" // var type is automatically decided by GO by the values assigned to Var

	fmt.Println(amount, unit, ", that's expensive...")
}
