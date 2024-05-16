package main

import "fmt"

type Pet struct { // struct named pet clubbing muliple variables
	name    string
	petType string
	age     int
}

func main() {
	//variable := structName{"name","PetType", age}
	nuggets := Pet{"Nuggets", "dog", 4} // this line is called a struct literal.
	mittens := Pet{"Mittens", "cat", 7} //struct literal creates a new instance of a struct
	robin := Pet{"Robin", "bird", 2}    // (in this case, Pet) with the specified values
	//fmt.Println(nuggets)
	//fmt.Println(mittens)
	//fmt.Println(robin)
	fmt.Println(nuggets, mittens, robin)
}
