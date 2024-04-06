package main

import (
	"fmt"
	"math"
)

// Define specialComputation() here
func specialComputation(x float64) float64 {

	calculate := math.Log2(math.Sqrt(math.Tan(x)))

	return calculate
}

func main() {
	var a, b, c, d float64
	a = .0214
	b = 1.02
	c = 0.312
	d = 4.001

	// Below the function is used multiple times which shows the reuse property of functions with multiple values.
	a = specialComputation(a)
	b = specialComputation(b)
	c = specialComputation(c)
	d = specialComputation(d)

	fmt.Println(a, b, c, d)
}
