package main

import "fmt"

type Point struct {
	x int
	y int
}

// Below struct named Circle is created
type Circle struct {
	//point  Point
	Point  // this is anonymous
	radius int
}

func main() {
	circle := Circle{Point{4, 5}, 2} // circle is instance of struct named Circle

	// Below the value of the x cordinated in the Point struct in Circle struct is extracted.
	//fmt.Println(circle.point.x)
	fmt.Println(circle.x) // when the Point field is now anonymous no need to give instance of struct.

	fmt.Println(circle)
}
