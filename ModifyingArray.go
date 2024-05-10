package main

import "fmt"

func main() {
	// I have 3 dogs, Frida, Fido, and Jeff but names are given wrong in the initial array which needs to be fixed with correct names
	myDogs := [3]string{"Frida", "Fedo", "Jegf"}
	//myDogs[1] = "Fido"
	//myDogs[2] = "Jeff"
	myDogs[1], myDogs[2] = "Fido", "Jeff"
	fmt.Println(myDogs)
}
