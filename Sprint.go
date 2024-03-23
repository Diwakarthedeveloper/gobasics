package main

import "fmt"

func main() {
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// oncatenate the strings step1 and step2:
	meditation := fmt.Sprintln(step1, step2)
	fmt.Print(meditation)

}
