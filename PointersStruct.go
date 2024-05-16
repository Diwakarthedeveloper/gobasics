package main

import "fmt"

type Rectangle struct {
	length float32
	width  float32
}

// Defining a function named area() below which does the computation for the rectangle,
// this function is also called method for structs with type float32
// Triangle is a struct created above and triangle is instance of a struct.
func (rect Rectangle) area() float32 {
	return rect.length * rect.width
}

func main() {

	rec := Rectangle{10, 5}
	fmt.Println(rec)

	fmt.Println(rec.area())

}
