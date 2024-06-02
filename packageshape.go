package main

import (
	"fmt"
	"shapes" // replace with the actual path to the shapes package wich should be in go installed location\Go\src\shapes
) // it should be insde source

func main() {
	radius := 5.0
	areaCircle := shapes.AreaOfCircle(radius)
	fmt.Printf("Area of Circle: %f\n", areaCircle)

	length := 4.0
	breadth := 8.0
	areaRectangle := shapes.AreaOfRectangle(length, breadth)
	fmt.Printf("Area of Rectangle: %f\n", areaRectangle)

	// This will result in an error as diameterOfCircle is unexported from shapes package
	// diameter := shapes.diameterOfCircle(radius)
	// fmt.Printf("Diameter of Circle: %f\n", diameter)
}
