package main

import "fmt"

func main() {
	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	//Assigning to wish the value of calling fmt.Sprintf() with the values template and pet.
	//wish should then contain the interpolated string "I wish I had a puppy.".

	wish = fmt.Sprintf(template, pet)

	fmt.Println(wish)
}
