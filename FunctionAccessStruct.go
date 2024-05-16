package main

import "fmt"

type Triangle struct { // struct named Triangle created here with variables height and base
	height float32
	base   float32
}

// Defining a function named area() below which does the computation for the rectangle,
// this function is also called method for structs with type float32
// Triangle is a struct created above and tri is instance of a struct.
func (tria Triangle) area() float32 {
	return 0.5 * (tria.base * tria.height)
}

func main() {

	tri := Triangle{10, 4} // here values are are passed in instance of a struct
	fmt.Println(tri)

	fmt.Println(tri.area())

}
