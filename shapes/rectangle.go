// rectangle.go in the shapes package
package shapes

// Exported Function - are once which can be clled in other packages , Name should start with Capital Letter
func AreaOfRectangle(length, breadth float64) float64 { // AreaOfRectangle can be easily called in other programs usinh shapes package
	return length * breadth
}

// unexported function - if name start with small letter it will
func perimeterOfRectangle(length, breadth float64) float64 {
	return 2 * (length + breadth)
}
